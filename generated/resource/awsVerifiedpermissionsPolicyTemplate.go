package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedpermissionsPolicyTemplate = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "statement": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedpermissionsPolicyTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedpermissionsPolicyTemplate), &result)
	return &result
}
