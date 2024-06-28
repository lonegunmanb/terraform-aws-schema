package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppfabricAppAuthorizationConnection = `{
  "block": {
    "attributes": {
      "app": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_authorization_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_bundle_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tenant": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "tenant_display_name": "string",
              "tenant_identifier": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "auth_request": {
        "block": {
          "attributes": {
            "code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "redirect_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppfabricAppAuthorizationConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppfabricAppAuthorizationConnection), &result)
	return &result
}
