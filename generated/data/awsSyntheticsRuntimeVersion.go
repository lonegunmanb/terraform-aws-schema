package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSyntheticsRuntimeVersion = `{
  "block": {
    "attributes": {
      "deprecation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "prefix": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSyntheticsRuntimeVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSyntheticsRuntimeVersion), &result)
	return &result
}
