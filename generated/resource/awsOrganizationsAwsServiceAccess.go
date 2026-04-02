package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOrganizationsAwsServiceAccess = `{
  "block": {
    "attributes": {
      "date_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_principal": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOrganizationsAwsServiceAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOrganizationsAwsServiceAccess), &result)
	return &result
}
