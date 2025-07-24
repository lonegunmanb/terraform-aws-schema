package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoLogDeliveryConfiguration = `{
  "block": {
    "attributes": {
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "log_configurations": {
        "block": {
          "attributes": {
            "event_source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_level": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloud_watch_logs_configuration": {
              "block": {
                "attributes": {
                  "log_group_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "firehose_configuration": {
              "block": {
                "attributes": {
                  "stream_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "s3_configuration": {
              "block": {
                "attributes": {
                  "bucket_arn": {
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

func AwsCognitoLogDeliveryConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoLogDeliveryConfiguration), &result)
	return &result
}
