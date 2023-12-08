package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminApplicationProviders = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "application_providers": {
        "block": {
          "attributes": {
            "application_provider_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "federation_protocol": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "display_data": {
              "block": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "display_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "icon_url": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsoadminApplicationProvidersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminApplicationProviders), &result)
	return &result
}
