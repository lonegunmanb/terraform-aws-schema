package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSpotDatafeedSubscription = `{
  "block": {
    "attributes": {
      "bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSpotDatafeedSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSpotDatafeedSubscription), &result)
	return &result
}
