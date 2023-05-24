package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketNotification = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventbridge": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "lambda_function": {
        "block": {
          "attributes": {
            "events": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "filter_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lambda_function_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "queue": {
        "block": {
          "attributes": {
            "events": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "filter_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "queue_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "topic": {
        "block": {
          "attributes": {
            "events": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "filter_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic_arn": {
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
  "version": 0
}`

func AwsS3BucketNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketNotification), &result)
	return &result
}
