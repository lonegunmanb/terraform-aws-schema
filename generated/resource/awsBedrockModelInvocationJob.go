package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockModelInvocationJob = `{
  "block": {
    "attributes": {
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "error_record_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_expiration_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_invocation_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "processed_record_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "skip_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "submit_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "success_record_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "timeout_duration_in_hours": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "total_record_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "input_data_config": {
        "block": {
          "block_types": {
            "s3_input_data_config": {
              "block": {
                "attributes": {
                  "s3_bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_input_format": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
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
      "output_data_config": {
        "block": {
          "block_types": {
            "s3_output_data_config": {
              "block": {
                "attributes": {
                  "s3_bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_encryption_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
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
  "version": 0
}`

func AwsBedrockModelInvocationJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockModelInvocationJob), &result)
	return &result
}
