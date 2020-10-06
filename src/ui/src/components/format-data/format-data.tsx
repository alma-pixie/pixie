import * as React from 'react';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import { formatBoolData, formatFloat64Data, formatUInt128Protobuf } from 'utils/format-data';
import {
  GaugeLevel,
  getCPULevel,
  getLatencyNSLevel,
} from 'utils/metric-thresholds';
import clsx from 'clsx';
import { DataType, SemanticType } from '../../types/generated/vizier_pb';

const JSON_INDENT_PX = 16;

export const AlertDataBase = ({ data, classes }) => (
  <div className={classes[`${data}`]}>{formatBoolData(data)}</div>
);

export const AlertData = withStyles((theme: Theme) => ({
  true: {
    color: theme.palette.error.dark,
  },
  false: {},
}))(AlertDataBase);

interface JSONDataProps {
  data: any;
  indentation?: number;
  multiline?: boolean;
  classes: any;
}

const jsonStyles = (theme: Theme) => createStyles({
  base: {
    fontFamily: '"Roboto Mono", serif',
    fontSize: '14px',
  },
  jsonKey: {
    color: theme.palette.foreground?.white,
  },
  number: {
    color: theme.palette.secondary.main,
  },
  null: {
    color: theme.palette.success.main,
  },
  string: {
    color: theme.palette.info.light,
    wordBreak: 'break-all',
  },
  boolean: {
    color: theme.palette.success.main,
  },
});

const JSONBase = React.memo<JSONDataProps>((props) => {
  const indentation = props.indentation ? props.indentation : 0;
  const { classes, multiline } = props;
  let { data } = props;
  let cls = String(typeof data);

  if (cls === 'string') {
    try {
      const parsedJson = JSON.parse(data);
      data = parsedJson;
    } catch {
      // Do nothing.
    }
  }

  if (data === null) {
    cls = 'null';
  }

  if (Array.isArray(data)) {
    return (
      <span className={classes.base}>
        {'[\u00A0'}
        {props.multiline ? <br /> : null}
        {
          data.map((val, idx) => (
            <span
              key={`${idx}-${indentation}`}
              style={{ marginLeft: multiline ? (indentation + 1) * JSON_INDENT_PX : 0 }}
            >
              <JSONData data={val} multiline={multiline} indentation={indentation + 1} />
              {idx !== Object.keys(data).length - 1 ? ',\u00A0' : ''}
              {multiline ? <br /> : null}
            </span>
          ))
        }
        <span style={{ marginLeft: multiline ? indentation * JSON_INDENT_PX : 0 }}>{'\u00A0]'}</span>
      </span>
    );
  }

  if (typeof data === 'object' && data !== null) {
    return (
      <span className={classes.base}>
        {'{\u00A0'}
        {props.multiline ? <br /> : null}
        {
          Object.keys(data).map((key, idx) => (
            <span
              key={`${key}-${indentation}`}
              style={{ marginLeft: props.multiline ? (indentation + 1) * JSON_INDENT_PX : 0 }}
            >
              <span className={classes.jsonKey}>
                {`${key}:\u00A0`}
              </span>
              <JSONData data={data[key]} multiline={props.multiline} indentation={indentation + 1} />
              {idx !== Object.keys(data).length - 1 ? ',\u00A0' : ''}
              {props.multiline ? <br /> : null}
            </span>
          ))
        }
        <span style={{ marginLeft: multiline ? indentation * JSON_INDENT_PX : 0 }}>{'\u00A0}'}</span>
      </span>
    );
  }
  return <span className={classes[cls]}>{String(data)}</span>;
});
// linter needs this for React.memo components.
JSONBase.displayName = 'JSONBase';

export const JSONData = withStyles(jsonStyles, {
  name: 'JSONData',
})(JSONBase);

interface GaugeDataProps {
  classes: any;
  data: any;
  getLevel: (number) => GaugeLevel;
}

const GaugeDataBase = ({ getLevel, classes, data }: GaugeDataProps) => {
  const floatVal = parseFloat(data);
  return (
    <div className={classes[getLevel(floatVal)]}>
      {formatFloat64Data(floatVal)}
    </div>
  );
};

const gaugeStyles = (theme: Theme) => createStyles({
  low: {
    color: theme.palette.success.dark,
  },
  med: {
    color: theme.palette.warning.dark,
  },
  high: {
    color: theme.palette.error.dark,
  },
});

export const GaugeData = withStyles(gaugeStyles, {
  name: 'GaugeData',
})(({ classes, data, level }: any) => (
  <div className={classes[level]}>
    {data}
  </div>
));

export const CPUData = withStyles(gaugeStyles, {
  name: 'CPUData',
})(({ classes, data }: any) => (
  <GaugeDataBase classes={classes} data={data} getLevel={getCPULevel} />
));

export const PortRendererBase = ({ data, classes }) => (
  <>
    <span className={classes.value}>{data}</span>
  </>
);

export const PortRenderer = withStyles(() => ({
  value: {
    fontFamily: '"Roboto Mono", serif',
    fontSize: '14px',
  },
}))(PortRendererBase);

export interface DataWithUnits {
  val: string;
  units: string;
}

export const formatScaled = (data: number, scale: number, suffixes: string[], decimals = 2): DataWithUnits => {
  const dm = decimals < 0 ? 0 : decimals;
  let i = Math.min(Math.floor(Math.log(Math.abs(data)) / Math.log(scale)), suffixes.length + 1);
  i = Math.max(0, Math.min(i, suffixes.length - 1));

  const val = `${parseFloat((data / (scale ** i)).toFixed(dm))}`;
  const units = suffixes[i];

  return {
    val,
    units,
  };
};

export const formatBytes = (data: number): DataWithUnits => (
  formatScaled(data,
    1024,
    ['\u00a0B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
    1)
);

export const formatDuration = (data: number): DataWithUnits => (
  formatScaled(data,
    1000,
    ['ns', '\u00b5s', 'ms', '\u00a0s'],
    1)
);

export const formatThroughput = (data: number): DataWithUnits => (
  formatScaled(data * 1E9,
    1000,
    ['\u00a0/s', 'K/s', 'M/s', 'B/s'],
    1)
);

export const formatThroughputBytes = (data: number): DataWithUnits => (
  formatScaled(data * 1E9,
    1024,
    ['\u00a0B/s', 'KB/s', 'MB/s', 'GB/s'],
    1)
);

const RenderValueWithUnitsBase = ({ data, classes }: { data: DataWithUnits; classes: any }) => (
  <>
    <span className={classes.value}>{`${data.val}\u00A0`}</span>
    <span className={classes.units}>{data.units}</span>
  </>
);

const RenderValueWithUnits = withStyles(() => ({
  units: {
    opacity: 0.5,
    fontFamily: '"Roboto Mono", serif',
    fontSize: '14px',
  },
  value: {
    fontFamily: '"Roboto Mono", serif',
    fontSize: '14px',
  },
}))(RenderValueWithUnitsBase);

export const BytesRenderer = ({ data }: { data: number }) => <RenderValueWithUnits data={formatBytes(data)} />;
export const DurationRenderer = ({ data }: { data: number }) => (
  <GaugeData
    data={<RenderValueWithUnits data={formatDuration(data)} />}
    level={getLatencyNSLevel(data)}
  />
);

const httpStatusCodeRendererStyles = (theme: Theme) => createStyles({
  root: {},
  unknown: {
    color: theme.palette.foreground.grey1,
  },
  oneHundredLevel: {
    color: theme.palette.success.main,
  },
  twoHundredLevel: {
    color: theme.palette.success.main,
  },
  threeHundredLevel: {
    color: theme.palette.success.main,
  },
  fourHundredLevel: {
    color: theme.palette.error.main,
  },
  fiveHundredLevel: {
    color: theme.palette.error.main,
  },
});

interface HTTPStatusCodeRendererProps extends WithStyles<typeof httpStatusCodeRendererStyles> {
  data: string;
}

export const HTTPStatusCodeRenderer = withStyles(httpStatusCodeRendererStyles)(
  ({ classes, data }: HTTPStatusCodeRendererProps) => {
    const intVal = parseInt(data, 10);
    const cls = clsx(
      classes.root,
      intVal < 0 && classes.unknown,
      intVal > 0 && intVal < 200 && classes.oneHundredLevel,
      intVal >= 200 && intVal < 300 && classes.twoHundredLevel,
      intVal >= 300 && intVal < 400 && classes.threeHundredLevel,
      intVal >= 400 && intVal < 500 && classes.fourHundredLevel,
      intVal >= 500 && classes.fiveHundredLevel);

    return (
      <>
        <span className={cls}>{data}</span>
      </>
    );
  });

export const formatPercent = (data: number): DataWithUnits => {
  const val = (100 * parseFloat(data.toString())).toFixed(1);
  const units = '%';

  return {
    val,
    units,
  };
};

interface PercentRendererProps {
  data: number;
}

export const PercentRenderer = ({ data }: PercentRendererProps) => <RenderValueWithUnits data={formatPercent(data)} />;

export const ThroughputRenderer = ({ data }: PercentRendererProps) => (
  <RenderValueWithUnits data={formatThroughput(data)} />
);

export const ThroughputBytesRenderer = ({ data }: PercentRendererProps) => (
  <RenderValueWithUnits data={formatThroughputBytes(data)} />
);

export const formatBySemType = (semType: SemanticType, val: any): DataWithUnits => {
  switch (semType) {
    case SemanticType.ST_BYTES:
      return formatBytes(val);
    case SemanticType.ST_DURATION_NS:
      return formatDuration(val);
    case SemanticType.ST_PERCENT:
      return formatPercent(val);
    case SemanticType.ST_THROUGHPUT_PER_NS:
      return formatThroughput(val);
    case SemanticType.ST_THROUGHPUT_BYTES_PER_NS:
      return formatThroughputBytes(val);
    default:
      break;
  }

  return {
    val,
    units: '',
  };
};

export const formatByDataType = (dataType: DataType, val: any): string => {
  switch (dataType) {
    case DataType.FLOAT64:
      return formatFloat64Data(val);
    case DataType.BOOLEAN:
      return formatBoolData(val);
    case DataType.UINT128:
      return formatUInt128Protobuf(val);
    case DataType.STRING:
    default:
      return val;
  }
};
