package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminApplicationAssignment = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "principal_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsoadminApplicationAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminApplicationAssignment), &result)
	return &result
}
