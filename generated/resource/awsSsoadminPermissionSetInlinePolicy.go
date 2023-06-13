package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminPermissionSetInlinePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inline_policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission_set_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsoadminPermissionSetInlinePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminPermissionSetInlinePolicy), &result)
	return &result
}
