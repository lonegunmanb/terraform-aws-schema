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
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
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
