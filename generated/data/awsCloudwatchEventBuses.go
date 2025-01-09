package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchEventBuses = `{
  "block": {
    "attributes": {
      "event_buses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "creation_time": "string",
              "description": "string",
              "last_modified_time": "string",
              "name": "string",
              "policy": "string"
            }
          ]
        ]
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchEventBusesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchEventBuses), &result)
	return &result
}
