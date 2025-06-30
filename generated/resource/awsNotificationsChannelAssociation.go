package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNotificationsChannelAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNotificationsChannelAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNotificationsChannelAssociation), &result)
	return &result
}
