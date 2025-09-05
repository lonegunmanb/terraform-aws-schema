package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoManagedLoginBranding = `{
  "block": {
    "attributes": {
      "client_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_login_branding_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "settings": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "settings_all": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_cognito_provided_values": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "asset": {
        "block": {
          "attributes": {
            "bytes": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "category": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "color_mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "extension": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoManagedLoginBrandingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoManagedLoginBranding), &result)
	return &result
}
