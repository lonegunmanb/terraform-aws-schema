package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMacie2OrganizationConfiguration = `{
  "block": {
    "attributes": {
      "auto_enable": {
        "description": "Whether to enable Amazon Macie automatically for accounts that are added to the organization in AWS Organizations",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMacie2OrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMacie2OrganizationConfiguration), &result)
	return &result
}
