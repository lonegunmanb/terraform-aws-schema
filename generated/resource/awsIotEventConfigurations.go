package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotEventConfigurations = `{
  "block": {
    "attributes": {
      "event_configurations": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "bool"
        ]
      },
      "id": {
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

func AwsIotEventConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotEventConfigurations), &result)
	return &result
}
