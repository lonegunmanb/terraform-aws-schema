package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentAgent = `{
  "block": {
    "attributes": {
      "agent_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "agent_collaboration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "agent_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_resource_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_encryption_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "foundation_model": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "guardrail_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "guardrail_identifier": "string",
              "guardrail_version": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "idle_session_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instruction": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "enabled_memory_types": [
                "list",
                "string"
              ],
              "session_summary_configuration": [
                "list",
                [
                  "object",
                  {
                    "max_recent_sessions": "number"
                  }
                ]
              ],
              "storage_days": "number"
            }
          ]
        ]
      },
      "prepare_agent": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "prepared_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_override_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "override_lambda": "string",
              "prompt_configurations": [
                "set",
                [
                  "object",
                  {
                    "base_prompt_template": "string",
                    "inference_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "max_length": "number",
                          "stop_sequences": [
                            "list",
                            "string"
                          ],
                          "temperature": "number",
                          "top_k": "number",
                          "top_p": "number"
                        }
                      ]
                    ],
                    "parser_mode": "string",
                    "prompt_creation_mode": "string",
                    "prompt_state": "string",
                    "prompt_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_resource_in_use_check": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockagentAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentAgent), &result)
	return &result
}
