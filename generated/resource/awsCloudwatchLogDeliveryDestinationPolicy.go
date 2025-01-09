package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchLogDeliveryDestinationPolicy = `{
  "block": {
    "attributes": {
      "delivery_destination_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delivery_destination_policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchLogDeliveryDestinationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchLogDeliveryDestinationPolicy), &result)
	return &result
}
