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

# -----------------------------------------------------------------------------
# Required
# -----------------------------------------------------------------------------

variable "dashboard_name" {
  description = "Name of the CloudWatch dashboard. Must be unique per account and region. Maximum 255 characters."
  type        = string

  validation {
    condition     = length(var.dashboard_name) >= 1 && length(var.dashboard_name) <= 255
    error_message = "Dashboard name must be between 1 and 255 characters."
  }
}

variable "dashboard_body" {
  description = "Detailed information about the dashboard, including widgets and layout, as a valid JSON string."
  type        = string

  validation {
    condition     = can(jsondecode(var.dashboard_body))
    error_message = "dashboard_body must be valid JSON."
  }
}

# -----------------------------------------------------------------------------
# Optional
# -----------------------------------------------------------------------------

variable "region" {
  description = "AWS region where the dashboard is managed. Defaults to the provider region when null."
  type        = string
  default     = null
}
