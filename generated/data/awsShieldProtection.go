package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsShieldProtection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protection_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protection_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsShieldProtectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsShieldProtection), &result)
	return &result
}
