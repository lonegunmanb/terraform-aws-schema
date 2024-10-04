package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamGroupPoliciesExclusive = `{
  "block": {
    "attributes": {
      "group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_names": {
        "description_kind": "plain",
        "required": true,
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

func AwsIamGroupPoliciesExclusiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamGroupPoliciesExclusive), &result)
	return &result
}
