package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogappregistryAttributeGroupAssociations = `{
  "block": {
    "attributes": {
      "attribute_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogappregistryAttributeGroupAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogappregistryAttributeGroupAssociations), &result)
	return &result
}
