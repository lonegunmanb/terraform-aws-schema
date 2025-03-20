package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreGroupMemberships = `{
  "block": {
    "attributes": {
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_memberships": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "group_id": "string",
              "identity_store_id": "string",
              "member_id": [
                "object",
                {
                  "user_id": "string"
                }
              ],
              "membership_id": "string"
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

func AwsIdentitystoreGroupMembershipsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreGroupMemberships), &result)
	return &result
}
