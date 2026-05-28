package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreOnlineEvaluationConfig = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_on_create": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "evaluation_execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "online_evaluation_config_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "online_evaluation_config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "online_evaluation_config_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloudwatch_config": [
                "list",
                [
                  "object",
                  {
                    "log_group_name": "string"
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
      "data_source_config": {
        "block": {
          "block_types": {
            "cloudwatch_logs": {
              "block": {
                "attributes": {
                  "log_group_names": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service_names": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "evaluator": {
        "block": {
          "attributes": {
            "evaluator_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rule": {
        "block": {
          "block_types": {
            "filter": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "value": {
                    "block": {
                      "attributes": {
                        "boolean_value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "double_value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "string_value": {
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
            "sampling_config": {
              "block": {
                "attributes": {
                  "sampling_percentage": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "session_config": {
              "block": {
                "attributes": {
                  "session_timeout_minutes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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

func AwsBedrockagentcoreOnlineEvaluationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreOnlineEvaluationConfig), &result)
	return &result
}
