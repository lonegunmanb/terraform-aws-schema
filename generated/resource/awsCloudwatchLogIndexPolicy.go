package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchLogIndexPolicy = `{
  "block": {
    "attributes": {
      "log_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "description": "Field index filter policy, in JSON",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchLogIndexPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchLogIndexPolicy), &result)
	return &result
}
