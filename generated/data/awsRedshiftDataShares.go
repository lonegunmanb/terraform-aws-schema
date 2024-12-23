package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftDataShares = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
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

func AwsRedshiftDataSharesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftDataShares), &result)
	return &result
}
