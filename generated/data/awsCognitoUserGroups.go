package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserGroups = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "groups": {
        "block": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "group_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "precedence": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoUserGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserGroups), &result)
	return &result
}
