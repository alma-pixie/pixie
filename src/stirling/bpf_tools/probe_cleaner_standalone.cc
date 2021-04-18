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

#include "src/common/base/base.h"
#include "src/stirling/bpf_tools/probe_cleaner.h"

DEFINE_string(cleanup_marker, ::px::stirling::utils::kPixieBPFProbeMarker,
              "Marker to search for when deleting probes.");

int main(int argc, char** argv) {
  px::EnvironmentGuard env_guard(&argc, argv);

  px::Status s = px::stirling::utils::CleanProbes(FLAGS_cleanup_marker);
  LOG_IF(ERROR, !s.ok()) << s.msg();

  return 0;
}
