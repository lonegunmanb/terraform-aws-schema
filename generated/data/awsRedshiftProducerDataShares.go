package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftProducerDataShares = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "producer_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_shares": {
        "block": {
          "attributes": {
            "data_share_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "managed_by": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "producer_arn": {
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
  "version": 0
}`

func AwsRedshiftProducerDataSharesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftProducerDataShares), &result)
	return &result
}
