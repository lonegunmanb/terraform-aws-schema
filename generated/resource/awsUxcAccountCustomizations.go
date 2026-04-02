package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsUxcAccountCustomizations = `{
  "block": {
    "attributes": {
      "account_color": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "visible_regions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "visible_services": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsUxcAccountCustomizationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsUxcAccountCustomizations), &result)
	return &result
}
