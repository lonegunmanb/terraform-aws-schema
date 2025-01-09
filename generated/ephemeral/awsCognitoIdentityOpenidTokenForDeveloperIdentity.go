package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoIdentityOpenidTokenForDeveloperIdentity = `{
  "block": {
    "attributes": {
      "identity_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logins": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "principal_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "token": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "token_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoIdentityOpenidTokenForDeveloperIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoIdentityOpenidTokenForDeveloperIdentity), &result)
	return &result
}
