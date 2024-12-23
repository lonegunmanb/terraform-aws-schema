package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentAgentVersions = `{
  "block": {
    "attributes": {
      "agent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "agent_version_summaries": {
        "block": {
          "attributes": {
            "agent_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "agent_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "agent_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "updated_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "guardrail_configuration": {
              "block": {
                "attributes": {
                  "guardrail_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "guardrail_version": {
                    "computed": true,
                    "description_kind": "plain",
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

func AwsBedrockagentAgentVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentAgentVersions), &result)
	return &result
}
