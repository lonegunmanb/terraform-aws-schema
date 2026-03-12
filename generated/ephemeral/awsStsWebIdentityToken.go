package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStsWebIdentityToken = `{
  "block": {
    "attributes": {
      "audience": {
        "description": "The intended recipients of the token (populates the ` + "`" + `aud` + "`" + ` claim in the JWT). Must contain between 1 and 10 items.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "duration_seconds": {
        "description": "The duration, in seconds, for which the JWT will remain valid. Value can range from 60 to 3600 seconds. Default is 300 seconds (5 minutes).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "expiration": {
        "computed": true,
        "description": "The expiration time of the token in RFC3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "signing_algorithm": {
        "description": "The cryptographic algorithm to use for signing the JWT. Valid values are ` + "`" + `RS256` + "`" + ` (RSA with SHA-256) and ` + "`" + `ES384` + "`" + ` (ECDSA using P-384 curve with SHA-384).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "web_identity_token": {
        "computed": true,
        "description": "The signed JWT token.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsStsWebIdentityTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsStsWebIdentityToken), &result)
	return &result
}
