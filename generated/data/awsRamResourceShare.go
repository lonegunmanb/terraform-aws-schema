package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRamResourceShare = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owning_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resource_owner": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_share_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
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

func AwsRamResourceShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRamResourceShare), &result)
	return &result
}
