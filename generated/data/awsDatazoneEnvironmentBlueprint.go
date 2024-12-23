package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatazoneEnvironmentBlueprint = `{
  "block": {
    "attributes": {
      "blueprint_provider": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "managed": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDatazoneEnvironmentBlueprintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatazoneEnvironmentBlueprint), &result)
	return &result
}
