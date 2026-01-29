package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNotificationsManagedNotificationAdditionalChannelAssociation = `{
  "block": {
    "attributes": {
      "channel_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_notification_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNotificationsManagedNotificationAdditionalChannelAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNotificationsManagedNotificationAdditionalChannelAssociation), &result)
	return &result
}
