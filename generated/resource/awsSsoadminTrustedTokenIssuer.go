package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminTrustedTokenIssuer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
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
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "trusted_token_issuer_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "trusted_token_issuer_configuration": {
        "block": {
          "block_types": {
            "oidc_jwt_configuration": {
              "block": {
                "attributes": {
                  "claim_attribute_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identity_store_attribute_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "issuer_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "jwks_retrieval_option": {
                    "description_kind": "plain",
                    "required": true,
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsoadminTrustedTokenIssuerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminTrustedTokenIssuer), &result)
	return &result
}
