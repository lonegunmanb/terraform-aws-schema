package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsService = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_rebalancing": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_provider_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "base": "number",
              "capacity_provider": "string",
              "weight": "number"
            }
          ]
        ]
      },
      "cluster_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alarms": [
                "list",
                [
                  "object",
                  {
                    "alarm_names": [
                      "list",
                      "string"
                    ],
                    "enable": "bool",
                    "rollback": "bool"
                  }
                ]
              ],
              "bake_time_in_minutes": "string",
              "canary_configuration": [
                "list",
                [
                  "object",
                  {
                    "canary_bake_time_in_minutes": "string",
                    "canary_percent": "number"
                  }
                ]
              ],
              "deployment_circuit_breaker": [
                "list",
                [
                  "object",
                  {
                    "enable": "bool",
                    "rollback": "bool"
                  }
                ]
              ],
              "lifecycle_hook": [
                "set",
                [
                  "object",
                  {
                    "hook_details": "string",
                    "hook_target_arn": "string",
                    "lifecycle_stages": [
                      "list",
                      "string"
                    ],
                    "role_arn": "string"
                  }
                ]
              ],
              "linear_configuration": [
                "list",
                [
                  "object",
                  {
                    "step_bake_time_in_minutes": "string",
                    "step_percent": "number"
                  }
                ]
              ],
              "maximum_percent": "number",
              "minimum_healthy_percent": "number",
              "strategy": "string"
            }
          ]
        ]
      },
      "deployment_controller": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "type": "string"
            }
          ]
        ]
      },
      "deployments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "desired_count": "number",
              "id": "string",
              "pending_count": "number",
              "running_count": "number",
              "status": "string",
              "task_definition": "string",
              "updated_at": "string"
            }
          ]
        ]
      },
      "desired_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "enable_ecs_managed_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_execute_command": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "events": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "id": "string",
              "message": "string"
            }
          ]
        ]
      },
      "health_check_grace_period_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "iam_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "advanced_configuration": [
                "list",
                [
                  "object",
                  {
                    "alternate_target_group_arn": "string",
                    "production_listener_rule": "string",
                    "role_arn": "string",
                    "test_listener_rule": "string"
                  }
                ]
              ],
              "container_name": "string",
              "container_port": "number",
              "elb_name": "string",
              "target_group_arn": "string"
            }
          ]
        ]
      },
      "network_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "assign_public_ip": "bool",
              "security_groups": [
                "set",
                "string"
              ],
              "subnets": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "ordered_placement_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "field": "string",
              "type": "string"
            }
          ]
        ]
      },
      "pending_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "placement_constraints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "expression": "string",
              "type": "string"
            }
          ]
        ]
      },
      "platform_family": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "running_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scheduling_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_registries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_name": "string",
              "container_port": "number",
              "port": "number",
              "registry_arn": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "task_definition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "created_at": "string",
              "id": "string",
              "pending_count": "number",
              "running_count": "number",
              "stability_status": "string",
              "status": "string",
              "task_definition": "string",
              "updated_at": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsService), &result)
	return &result
}
