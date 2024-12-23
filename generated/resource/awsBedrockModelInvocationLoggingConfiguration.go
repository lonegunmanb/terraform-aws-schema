package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockModelInvocationLoggingConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "logging_config": {
        "block": {
          "attributes": {
            "embedding_data_delivery_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "image_data_delivery_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "text_data_delivery_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "cloudwatch_config": {
              "block": {
                "attributes": {
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "large_data_delivery_s3_config": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_prefix": {
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
              "nesting_mode": "single"
            },
            "s3_config": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_prefix": {
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
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockModelInvocationLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockModelInvocationLoggingConfiguration), &result)
	return &result
}
