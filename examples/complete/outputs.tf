// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

output "region" {
  description = "The AWS region where resources are deployed."
  value       = coalesce(var.region, data.aws_region.current.region)
}

output "id" {
  description = "The ID of the dashboard (same as the name)."
  value       = module.dashboard.id
}

output "arn" {
  description = "The ARN of the dashboard."
  value       = module.dashboard.arn
}

output "name" {
  description = "The name of the dashboard."
  value       = module.dashboard.name
}

output "dashboard_body" {
  description = "The dashboard body JSON configured by the module."
  value       = module.dashboard.dashboard_body
}
