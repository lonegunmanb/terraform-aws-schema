package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrAccountSetting = `{
  "block": {
    "attributes": {
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "value": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrAccountSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrAccountSetting), &result)
	return &result
}
