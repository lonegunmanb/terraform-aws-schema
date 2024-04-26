package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "display_name": "string",
              "external_ids": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "issuer": "string"
                  }
                ]
              ],
              "group_id": "string",
              "identity_store_id": "string"
            }
          ]
        ]
      },
      "identity_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIdentitystoreGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreGroups), &result)
	return &result
}
