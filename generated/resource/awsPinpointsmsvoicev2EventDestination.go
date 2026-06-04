package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointsmsvoicev2EventDestination = `{
  "block": {
    "attributes": {
      "configuration_set_arn": {
        "computed": true,
        "description": "ARN of the parent configuration set.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_set_name": {
        "description": "Name of the configuration set this event destination belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Whether the event destination is enabled. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_destination_name": {
        "description": "Name of the event destination.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "matching_event_types": {
        "description": "Event types for which the destination receives records.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
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
      "cloudwatch_logs_destination": {
        "block": {
          "attributes": {
            "iam_role_arn": {
              "description": "ARN of the IAM role that End User Messaging SMS assumes to write to the log group.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_group_arn": {
              "description": "ARN of the Amazon CloudWatch log group that receives the events.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "kinesis_firehose_destination": {
        "block": {
          "attributes": {
            "delivery_stream_arn": {
              "description": "ARN of the Amazon Data Firehose delivery stream that receives the events.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "iam_role_arn": {
              "description": "ARN of the IAM role that End User Messaging SMS assumes to write to the delivery stream.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sns_destination": {
        "block": {
          "attributes": {
            "topic_arn": {
              "description": "ARN of the Amazon SNS topic that receives the events.",
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

func AwsPinpointsmsvoicev2EventDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointsmsvoicev2EventDestination), &result)
	return &result
}
