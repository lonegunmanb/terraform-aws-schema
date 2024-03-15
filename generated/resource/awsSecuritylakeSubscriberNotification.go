package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecuritylakeSubscriberNotification = `{
  "block": {
    "attributes": {
      "endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "block_types": {
            "https_notification_configuration": {
              "block": {
                "attributes": {
                  "authorization_api_key_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "authorization_api_key_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "sqs_notification_configuration": {
              "block": {
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

func AwsSecuritylakeSubscriberNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecuritylakeSubscriberNotification), &result)
	return &result
}
