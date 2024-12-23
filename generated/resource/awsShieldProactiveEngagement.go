package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsShieldProactiveEngagement = `{
  "block": {
    "attributes": {
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "emergency_contact": {
        "block": {
          "attributes": {
            "contact_notes": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "phone_number": {
              "description_kind": "plain",
              "optional": true,
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

func AwsShieldProactiveEngagementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsShieldProactiveEngagement), &result)
	return &result
}
