package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDevopsguruResourceCollection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "cloudformation": {
        "block": {
          "attributes": {
            "stack_names": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "tags": {
        "block": {
          "attributes": {
            "app_boundary_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tag_values": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
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

func AwsDevopsguruResourceCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDevopsguruResourceCollection), &result)
	return &result
}
