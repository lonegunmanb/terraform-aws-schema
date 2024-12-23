package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheReservedCacheNodeOffering = `{
  "block": {
    "attributes": {
      "cache_node_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fixed_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "offering_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "offering_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheReservedCacheNodeOfferingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheReservedCacheNodeOffering), &result)
	return &result
}
