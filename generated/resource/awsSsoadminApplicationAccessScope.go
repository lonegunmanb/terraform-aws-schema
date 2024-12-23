package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminApplicationAccessScope = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorized_targets": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsoadminApplicationAccessScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminApplicationAccessScope), &result)
	return &result
}
