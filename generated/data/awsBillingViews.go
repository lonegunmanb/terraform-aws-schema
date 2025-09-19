package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBillingViews = `{
  "block": {
    "attributes": {
      "billing_view": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "billing_view_type": "string",
              "description": "string",
              "name": "string",
              "owner_account_id": "string"
            }
          ]
        ]
      },
      "billing_view_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBillingViewsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBillingViews), &result)
	return &result
}
