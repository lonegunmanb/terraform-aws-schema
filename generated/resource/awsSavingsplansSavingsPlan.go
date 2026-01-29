package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSavingsplansSavingsPlan = `{
  "block": {
    "attributes": {
      "commitment": {
        "description": "The hourly commitment, in USD.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "currency": {
        "computed": true,
        "description": "The currency of the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description.",
        "description_kind": "plain",
        "type": "string"
      },
      "ec2_instance_family": {
        "computed": true,
        "description": "The EC2 instance family for the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "end": {
        "computed": true,
        "description": "The end time of the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "offering_id": {
        "computed": true,
        "description": "The ID of the offering.",
        "description_kind": "plain",
        "type": "string"
      },
      "payment_option": {
        "computed": true,
        "description": "The payment option for the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_types": {
        "computed": true,
        "description": "The product types.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "purchase_time": {
        "description": "The time at which to purchase the Savings Plan, in UTC format (YYYY-MM-DDTHH:MM:SSZ).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recurring_payment_amount": {
        "computed": true,
        "description": "The recurring payment amount.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "returnable_until": {
        "computed": true,
        "description": "The recurring payment amount.",
        "description_kind": "plain",
        "type": "string"
      },
      "savings_plan_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "savings_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "savings_plan_offering_id": {
        "description": "The unique ID of a Savings Plan offering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "savings_plan_type": {
        "computed": true,
        "description": "The type of Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "start": {
        "computed": true,
        "description": "The start time of the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the Savings Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "term_duration_in_seconds": {
        "computed": true,
        "description": "The duration of the term, in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "upfront_payment_amount": {
        "description": "The up-front payment amount.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
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

func AwsSavingsplansSavingsPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSavingsplansSavingsPlan), &result)
	return &result
}
