package data

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
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "lambda_function": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "events": [
                "set",
                "string"
              ],
              "filter_prefix": "string",
              "filter_suffix": "string",
              "id": "string",
              "lambda_function_arn": "string"
            }
          ]
        ]
      },
      "queue": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "events": [
                "set",
                "string"
              ],
              "filter_prefix": "string",
              "filter_suffix": "string",
              "id": "string",
              "queue_arn": "string"
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
      "topic": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "events": [
                "set",
                "string"
              ],
              "filter_prefix": "string",
              "filter_suffix": "string",
              "id": "string",
              "topic_arn": "string"
            }
          ]
        ]
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
