package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsUxcServices = `{
  "block": {
    "attributes": {
      "services": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsUxcServicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsUxcServices), &result)
	return &result
}
