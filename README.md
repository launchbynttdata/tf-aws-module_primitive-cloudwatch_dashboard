# TF AWS Module Primitive - CloudWatch Dashboard

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module creates an [AWS CloudWatch dashboard](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_dashboard) for visualizing metrics, logs, and alarms.

## Pre-Commit Hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) defines pre-commit hooks for Terraform, Go, and common linting. The `commitlint` hook enforces conventional commit format. The `detect-secrets-hook` prevents new secrets from being introduced. See [pre-commit](https://pre-commit.com/#install) for installation. Install the commit-msg hook manually:

```
pre-commit install --hook-type commit-msg
```

## Usage

See [examples/complete](examples/complete) for a full working example.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.10 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.100, < 7.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_dashboard.dashboard](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_dashboard) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_dashboard_body"></a> [dashboard\_body](#input\_dashboard\_body) | Detailed information about the dashboard, including widgets and layout, as a JSON string. | `string` | n/a | yes |
| <a name="input_dashboard_name"></a> [dashboard\_name](#input\_dashboard\_name) | Name of the CloudWatch dashboard. Must be unique per account and region. Maximum 255 characters. | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | AWS region where the dashboard is managed. Defaults to the provider region when null. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the dashboard. |
| <a name="output_dashboard_body"></a> [dashboard\_body](#output\_dashboard\_body) | The dashboard body JSON. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the dashboard (same as the name). |
| <a name="output_name"></a> [name](#output\_name) | The name of the dashboard. |
<!-- END_TF_DOCS -->
