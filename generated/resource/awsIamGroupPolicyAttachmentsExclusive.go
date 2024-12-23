package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamGroupPolicyAttachmentsExclusive = `{
  "block": {
    "attributes": {
      "group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_arns": {
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

func AwsIamGroupPolicyAttachmentsExclusiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamGroupPolicyAttachmentsExclusive), &result)
	return &result
}
