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
      "prepare_agent": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      }
    },
    "block_types": {
      "action_group_executor": {
        "block": {
          "attributes": {
            "custom_control": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
      },
      "function_schema": {
        "block": {
          "block_types": {
            "member_functions": {
              "block": {
                "block_types": {
                  "functions": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "description": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "required": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
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
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsBedrockagentAgentActionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentAgentActionGroup), &result)
	return &result
}
