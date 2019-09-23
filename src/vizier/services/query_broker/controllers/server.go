package controllers

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pixielabs.ai/pixielabs/src/carnot/compiler"
	"pixielabs.ai/pixielabs/src/carnot/compiler/compilerpb"
	"pixielabs.ai/pixielabs/src/carnot/compiler/distributedpb"
	planpb "pixielabs.ai/pixielabs/src/carnot/planpb"
	statuspb "pixielabs.ai/pixielabs/src/common/base/proto"
	schemapb "pixielabs.ai/pixielabs/src/table_store/proto"
	"pixielabs.ai/pixielabs/src/utils"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/metadatapb"
	"pixielabs.ai/pixielabs/src/vizier/services/query_broker/querybrokerenv"
	"pixielabs.ai/pixielabs/src/vizier/services/query_broker/querybrokerpb"
)

// Planner describes the interface for any planner.
type Planner interface {
	Plan(planState *distributedpb.LogicalPlannerState, query string) (*distributedpb.LogicalPlannerResult, error)
	Free()
}

// Executor is the interface for a query executor.
type Executor interface {
	ExecuteQuery(planMap map[uuid.UUID]*planpb.Plan) error
	WaitForCompletion() ([]*querybrokerpb.VizierQueryResponse_ResponseByAgent, error)
	AddResult(res *querybrokerpb.AgentQueryResultRequest)
	GetQueryID() uuid.UUID
}

// Server defines an gRPC server type.
type Server struct {
	env         querybrokerenv.QueryBrokerEnv
	mdsClient   metadatapb.MetadataServiceClient
	natsConn    *nats.Conn
	newExecutor func(*nats.Conn, uuid.UUID, *[]uuid.UUID) Executor
	planner     Planner
	executors   map[uuid.UUID]Executor
	// Mutex is used for managing query executor instances.
	mux sync.Mutex
}

// NewServer creates GRPC handlers.
func NewServer(env querybrokerenv.QueryBrokerEnv, mdsClient metadatapb.MetadataServiceClient, natsConn *nats.Conn, planner Planner) (*Server, error) {
	return newServer(env, mdsClient, natsConn, NewQueryExecutor, planner)
}

func newServer(env querybrokerenv.QueryBrokerEnv,
	mdsClient metadatapb.MetadataServiceClient,
	natsConn *nats.Conn,
	newExecutor func(*nats.Conn, uuid.UUID, *[]uuid.UUID) Executor,
	planner Planner) (*Server, error) {

	return &Server{
		env:         env,
		mdsClient:   mdsClient,
		natsConn:    natsConn,
		newExecutor: newExecutor,
		planner:     planner,
		executors:   make(map[uuid.UUID]Executor),
	}, nil
}

func (s *Server) trackExecutorForQuery(executor Executor) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.executors[executor.GetQueryID()] = executor
}

func (s *Server) deleteExecutorForQuery(queryID uuid.UUID) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.executors, queryID)
}

func failedStatusQueryResponse(status *statuspb.Status) *querybrokerpb.VizierQueryResponse {
	queryResponse := &querybrokerpb.VizierQueryResponse{
		Status: status,
	}
	return queryResponse
}

func formatCompilerError(status *statuspb.Status) (string, error) {
	var errorPB compilerpb.CompilerErrorGroup
	err := compiler.GetCompilerErrorContext(status, &errorPB)
	if err != nil {
		return "", err
	}
	return proto.MarshalTextString(&errorPB), nil
}

func makeAgentCarnotInfo(agentID uuid.UUID) *distributedpb.CarnotInfo {
	// TODO(philkuz) (PL-910) need to update this to also contain table info.
	return &distributedpb.CarnotInfo{
		QueryBrokerAddress:   agentID.String(),
		HasGrpcServer:        false,
		HasDataStore:         true,
		ProcessesData:        true,
		AcceptsRemoteSources: false,
	}
}

func makeKelvinCarnotInfo(agentID uuid.UUID, grpcAddress string) *distributedpb.CarnotInfo {
	return &distributedpb.CarnotInfo{
		QueryBrokerAddress:   agentID.String(),
		HasGrpcServer:        true,
		GrpcAddress:          grpcAddress,
		HasDataStore:         false,
		ProcessesData:        true,
		AcceptsRemoteSources: false,
	}
}

func makePlannerState(agentList []uuid.UUID, kelvinList []uuid.UUID, schema *schemapb.Schema) (*distributedpb.LogicalPlannerState, error) {
	// TODO(philkuz) (PL-910) need to update this to pass table info.
	carnotInfoList := make([]*distributedpb.CarnotInfo, 0)
	for _, agentID := range agentList {
		carnotInfoList = append(carnotInfoList, makeAgentCarnotInfo(agentID))
	}

	for _, kelvinID := range kelvinList {
		// TODO(philkuz/zasgar) (PL-873) pass the grpc address for kelvin here somehow.
		// Maybe change KelvinLIst to be a list of structs that contain uuid and grpc address instead?
		kelvinGrpcAddress := ""
		carnotInfoList = append(carnotInfoList, makeKelvinCarnotInfo(kelvinID, kelvinGrpcAddress))
	}

	plannerState := distributedpb.LogicalPlannerState{
		Schema: schema,
		DistributedState: &distributedpb.DistributedState{
			CarnotInfo: carnotInfoList,
		},
	}
	return &plannerState, nil
}

// ExecuteQuery executes a query on multiple agents and compute node.
func (s *Server) ExecuteQuery(ctx context.Context, req *querybrokerpb.QueryRequest) (*querybrokerpb.VizierQueryResponse, error) {
	// Get the table schema that is presumably shared across agents.
	mdsSchemaReq := &metadatapb.SchemaRequest{}
	mdsSchemaResp, err := s.mdsClient.GetSchemas(ctx, mdsSchemaReq)
	if err != nil {
		log.WithError(err).Fatal("Failed to get schemas.")
		return nil, err
	}
	schema := mdsSchemaResp.Schema

	// Get all available agents for now.
	mdsReq := &metadatapb.AgentInfoRequest{}
	mdsResp, err := s.mdsClient.GetAgentInfo(ctx, mdsReq)
	if err != nil {
		return nil, err
	}

	var agentList []uuid.UUID

	for _, info := range mdsResp.Info {
		agentIDPB := info.Info.AgentID
		agentID, err := utils.UUIDFromProto(agentIDPB)
		if err != nil {
			return nil, err
		}
		agentList = append(agentList, agentID)
	}

	// TODO(philkuz/zasgar) (PL-873) populate kelvinList and pass the grpc address for kelvin here somehow.
	// Maybe change KelvinLIst to be a list of structs that contain uuid and grpc address instead?
	var kelvinList []uuid.UUID

	plannerState, err := makePlannerState(agentList, kelvinList, schema)
	if err != nil {
		return nil, err
	}

	// Compile the query plan.
	plannerResultPB, err := s.planner.Plan(plannerState, req.QueryStr)
	if err != nil {
		return nil, err
	}

	// TODO(philkuz/michelle) do we want to return the problem through the error
	// or through the message? I've done both for now because the error route is
	// the only one that appears in the front end.
	if plannerResultPB.Status.ErrCode != statuspb.OK {
		errorStr, err := formatCompilerError(plannerResultPB.Status)
		if err != nil {
			log.WithError(err).Errorf("Couldn't get the compiler status for message: %s", proto.MarshalTextString(plannerResultPB.Status))
			return nil, fmt.Errorf("Error occurred without line and column: '%s'", plannerResultPB.Status.Msg)
		}
		return failedStatusQueryResponse(plannerResultPB.Status), fmt.Errorf(errorStr)
	}

	// Plan describes the mapping of agents to the plan that should execute on them.
	plan := plannerResultPB.Plan
	planMap := make(map[uuid.UUID]*planpb.Plan)

	for carnotID, agentPlan := range plan.QbAddressToPlan {
		u, err := uuid.FromString(carnotID)
		if err != nil {
			log.WithError(err).Fatalf("Couldn't parse uuid from agent id \"%s\"", carnotID)
			return nil, err
		}
		planMap[u] = agentPlan
	}

	planCarnotList := make([]uuid.UUID, 0, len(planMap))
	for carnotID := range planMap {
		planCarnotList = append(planCarnotList, carnotID)
	}

	queryID := uuid.NewV4()
	queryExecutor := s.newExecutor(s.natsConn, queryID, &planCarnotList)

	s.trackExecutorForQuery(queryExecutor)

	if err := queryExecutor.ExecuteQuery(planMap); err != nil {
		return nil, err
	}

	responses, err := queryExecutor.WaitForCompletion()
	if err != nil {
		return nil, err
	}

	queryResponse := &querybrokerpb.VizierQueryResponse{
		Responses: responses,
	}

	s.deleteExecutorForQuery(queryID)

	return queryResponse, nil
}

// GetSchemas returns the schemas in the system.
func (s *Server) GetSchemas(ctx context.Context, req *querybrokerpb.SchemaRequest) (*querybrokerpb.SchemaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented yet")
}

// GetAgentInfo returns information about registered agents.
func (s *Server) GetAgentInfo(ctx context.Context, req *querybrokerpb.AgentInfoRequest) (*querybrokerpb.AgentInfoResponse, error) {
	mdsReq := &metadatapb.AgentInfoRequest{}
	mdsResp, err := s.mdsClient.GetAgentInfo(ctx, mdsReq)
	if err != nil {
		return nil, err
	}

	var agentStatuses []*querybrokerpb.AgentStatus

	for _, agentInfo := range mdsResp.Info {
		agentStatuses = append(agentStatuses, &querybrokerpb.AgentStatus{
			Info: &querybrokerpb.AgentInfo{
				AgentID: agentInfo.Info.AgentID,
				HostInfo: &querybrokerpb.HostInfo{
					Hostname: agentInfo.Info.HostInfo.Hostname,
				},
			},
			LastHeartbeatNs: agentInfo.LastHeartbeatNs,
			State:           querybrokerpb.AGENT_STATE_HEALTHY,
		})
	}

	// Map the responses.
	resp := &querybrokerpb.AgentInfoResponse{
		Info: agentStatuses,
	}

	return resp, nil
}

// ReceiveAgentQueryResult gets the query result from an agent and stores the results until all
// relevant agents have responded.
func (s *Server) ReceiveAgentQueryResult(ctx context.Context, req *querybrokerpb.AgentQueryResultRequest) (*querybrokerpb.AgentQueryResultResponse, error) {
	queryIDPB := req.Result.QueryID
	queryID, err := utils.UUIDFromProto(queryIDPB)
	if err != nil {
		return nil, err
	}

	setExecutor := func(queryID uuid.UUID) (Executor, bool) {
		s.mux.Lock()
		exec, ok := s.executors[queryID]
		defer s.mux.Unlock()
		return exec, ok
	}

	exec, ok := setExecutor(queryID)

	if !ok {
		return nil, errors.New("Query ID not present")
	}

	exec.AddResult(req)

	queryResponse := &querybrokerpb.AgentQueryResultResponse{}

	return queryResponse, nil
}
