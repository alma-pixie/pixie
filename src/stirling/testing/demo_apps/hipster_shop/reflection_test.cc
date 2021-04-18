/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include "src/stirling/testing/demo_apps/hipster_shop/reflection.h"

#include <google/protobuf/descriptor.pb.h>

#include "src/common/testing/testing.h"

namespace demos {
namespace hipster_shop {

using ::google::protobuf::FileDescriptorSet;
using ::px::testing::proto::EqualsProto;
using ::px::testing::proto::Partially;
using ::testing::ElementsAre;

TEST(GetFileDescriptorSetTest, HasAllServicesAndMessages) {
  FileDescriptorSet fds = GetFileDescriptorSet();
  EXPECT_THAT(fds.file(), ElementsAre(Partially(EqualsProto(R"proto(
      syntax: "proto3"
      name: "src/stirling/testing/demo_apps/hipster_shop/proto/demo.proto"
      package: "hipstershop"
  )proto"))));
}

}  // namespace hipster_shop
}  // namespace demos
