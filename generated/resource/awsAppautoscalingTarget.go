package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppautoscalingTarget = `{
  "block": {
    "attributes": {
      "arn": {
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
      "max_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scalable_dimension": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "suspended_state": {
        "block": {
          "attributes": {
            "dynamic_scaling_in_suspended": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dynamic_scaling_out_suspended": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scheduled_scaling_suspended": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AwsAppautoscalingTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppautoscalingTarget), &result)
	return &result
}
