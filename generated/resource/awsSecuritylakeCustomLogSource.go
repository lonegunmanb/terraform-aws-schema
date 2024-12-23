package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecuritylakeCustomLogSource = `{
  "block": {
    "attributes": {
      "attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "crawler_arn": "string",
              "database_arn": "string",
              "table_arn": "string"
            }
          ]
        ]
      },
      "event_classes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_details": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "location": "string",
              "role_arn": "string"
            }
          ]
        ]
      },
      "source_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "block_types": {
            "crawler_configuration": {
              "block": {
                "attributes": {
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "provider_identity": {
              "block": {
                "attributes": {
                  "external_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "principal": {
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

func AwsSecuritylakeCustomLogSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecuritylakeCustomLogSource), &result)
	return &result
}
