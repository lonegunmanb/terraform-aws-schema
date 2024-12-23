package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubStandardsControlAssociations = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_control_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "standards_control_associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "association_status": "string",
              "related_requirements": [
                "list",
                "string"
              ],
              "security_control_arn": "string",
              "security_control_id": "string",
              "standards_arn": "string",
              "standards_control_description": "string",
              "standards_control_title": "string",
              "updated_at": "string",
              "updated_reason": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubStandardsControlAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubStandardsControlAssociations), &result)
	return &result
}
