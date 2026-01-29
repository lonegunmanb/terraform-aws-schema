package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNotificationsOrganizationalUnitAssociation = `{
  "block": {
    "attributes": {
      "notification_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organizational_unit_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNotificationsOrganizationalUnitAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNotificationsOrganizationalUnitAssociation), &result)
	return &result
}
