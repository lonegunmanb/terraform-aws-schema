package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmPatchBaselines = `{
  "block": {
    "attributes": {
      "baseline_identities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "baseline_description": "string",
              "baseline_id": "string",
              "baseline_name": "string",
              "default_baseline": "bool",
              "operating_system": "string"
            }
          ]
        ]
      },
      "default_baselines": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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

func AwsSsmPatchBaselinesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmPatchBaselines), &result)
	return &result
}
