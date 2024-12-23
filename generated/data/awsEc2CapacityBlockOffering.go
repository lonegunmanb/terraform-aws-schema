package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2CapacityBlockOffering = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_block_offering_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_duration_hours": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "currency_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_date_range": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_date_range": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upfront_fee": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2CapacityBlockOfferingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2CapacityBlockOffering), &result)
	return &result
}
