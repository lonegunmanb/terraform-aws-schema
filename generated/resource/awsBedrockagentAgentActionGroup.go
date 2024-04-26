package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentAgentActionGroup = `{
  "block": {
    "attributes": {
      "action_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "action_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "action_group_state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent_action_group_signature": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_resource_in_use_check": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "action_group_executor": {
        "block": {
          "attributes": {
            "lambda": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "api_schema": {
        "block": {
          "attributes": {
            "payload": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "s3": {
              "block": {
                "attributes": {
                  "s3_bucket_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_object_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockagentAgentActionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentAgentActionGroup), &result)
	return &result
}
