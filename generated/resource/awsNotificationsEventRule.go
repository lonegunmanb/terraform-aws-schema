package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNotificationsEventRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_pattern": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "regions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNotificationsEventRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNotificationsEventRule), &result)
	return &result
}
