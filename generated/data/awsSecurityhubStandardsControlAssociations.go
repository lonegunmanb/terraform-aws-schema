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
      }
    },
    "block_types": {
      "standards_control_associations": {
        "block": {
          "attributes": {
            "association_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "related_requirements": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "security_control_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "security_control_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "standards_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "standards_control_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "standards_control_title": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "updated_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "updated_reason": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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
