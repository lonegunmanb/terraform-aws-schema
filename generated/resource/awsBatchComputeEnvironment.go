package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBatchComputeEnvironment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_environment_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_environment_name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecc_cluster_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ecs_cluster_arn": {
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
      "service_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "compute_resources": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bid_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "desired_vcpus": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ec2_key_pair": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "max_vcpus": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_vcpus": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "spot_iam_fleet_role": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "launch_template": {
              "block": {
                "attributes": {
                  "launch_template_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_template_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBatchComputeEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBatchComputeEnvironment), &result)
	return &result
}