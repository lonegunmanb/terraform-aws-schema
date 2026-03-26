package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSavingsplansOfferings = `{
  "block": {
    "attributes": {
      "currencies": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "descriptions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "durations": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "offering_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "offerings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "currency": "string",
              "description": "string",
              "duration_seconds": "number",
              "offering_id": "string",
              "operation": "string",
              "payment_option": "string",
              "plan_type": "string",
              "product_types": [
                "set",
                "string"
              ],
              "properties": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "service_code": "string",
              "usage_type": "string"
            }
          ]
        ]
      },
      "operations": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "payment_options": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "plan_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "product_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_codes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "usage_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSavingsplansOfferingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSavingsplansOfferings), &result)
	return &result
}
