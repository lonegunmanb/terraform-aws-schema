package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdawebFunctionRevision = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "build_config": {
        "block": {
          "block_types": {
            "code_config": {
              "block": {
                "block_types": {
                  "s3_object": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version_id": {
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
            "runtime_config": {
              "block": {
                "attributes": {
                  "runtime": {
                    "description_kind": "plain",
                    "required": true,
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
      "service_config": {
        "block": {
          "attributes": {
            "egress_network_connector_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            },
            "execution_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_concurrency_per_environment": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "telemetry_config_actual": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "logging_config": [
                      "list",
                      [
                        "object",
                        {
                          "application_log_level": "string",
                          "log_group": "string",
                          "system_log_level": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            },
            "timeout_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "telemetry_config": {
              "block": {
                "block_types": {
                  "logging_config": {
                    "block": {
                      "attributes": {
                        "application_log_level": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "log_group": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "system_log_level": {
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

func AwsLambdawebFunctionRevisionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdawebFunctionRevision), &result)
	return &result
}
