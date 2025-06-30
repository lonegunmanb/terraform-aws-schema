package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOrganizationsOrganizationalUnitDescendantOrganizationalUnits = `{
  "block": {
    "attributes": {
      "children": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOrganizationsOrganizationalUnitDescendantOrganizationalUnitsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOrganizationsOrganizationalUnitDescendantOrganizationalUnits), &result)
	return &result
}
