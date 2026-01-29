package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNotificationsManagedNotificationAccountContactAssociation = `{
  "block": {
    "attributes": {
      "contact_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_notification_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNotificationsManagedNotificationAccountContactAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNotificationsManagedNotificationAccountContactAssociation), &result)
	return &result
}
