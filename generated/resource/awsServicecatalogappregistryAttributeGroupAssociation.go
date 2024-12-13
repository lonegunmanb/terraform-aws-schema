package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogappregistryAttributeGroupAssociation = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "ID of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "attribute_group_id": {
        "description": "ID of the attribute group to associate with the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogappregistryAttributeGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogappregistryAttributeGroupAssociation), &result)
	return &result
}
