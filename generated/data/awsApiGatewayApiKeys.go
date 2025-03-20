package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayApiKeys = `{
  "block": {
    "attributes": {
      "customer_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "include_values": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "items": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_date": "string",
              "customer_id": "string",
              "description": "string",
              "enabled": "bool",
              "id": "string",
              "last_updated_date": "string",
              "name": "string",
              "stage_keys": [
                "list",
                "string"
              ],
              "tags": [
                "map",
                "string"
              ],
              "value": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApiGatewayApiKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayApiKeys), &result)
	return &result
}
