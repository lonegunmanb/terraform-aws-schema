package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOrganizationsEntityPath = `{
  "block": {
    "attributes": {
      "entity_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOrganizationsEntityPathSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOrganizationsEntityPath), &result)
	return &result
}
