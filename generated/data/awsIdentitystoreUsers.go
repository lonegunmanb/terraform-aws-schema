package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreUsers = `{
  "block": {
    "attributes": {
      "identity_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "addresses": [
                "list",
                [
                  "object",
                  {
                    "country": "string",
                    "formatted": "string",
                    "locality": "string",
                    "postal_code": "string",
                    "primary": "bool",
                    "region": "string",
                    "street_address": "string",
                    "type": "string"
                  }
                ]
              ],
              "display_name": "string",
              "emails": [
                "list",
                [
                  "object",
                  {
                    "primary": "bool",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
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
              "identity_store_id": "string",
              "locale": "string",
              "name": [
                "list",
                [
                  "object",
                  {
                    "family_name": "string",
                    "formatted": "string",
                    "given_name": "string",
                    "honorific_prefix": "string",
                    "honorific_suffix": "string",
                    "middle_name": "string"
                  }
                ]
              ],
              "nickname": "string",
              "phone_numbers": [
                "list",
                [
                  "object",
                  {
                    "primary": "bool",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "preferred_language": "string",
              "profile_url": "string",
              "timezone": "string",
              "title": "string",
              "user_id": "string",
              "user_name": "string",
              "user_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIdentitystoreUsersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreUsers), &result)
	return &result
}
