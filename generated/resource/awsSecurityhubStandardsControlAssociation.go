package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubStandardsControlAssociation = `{
  "block": {
    "attributes": {
      "association_status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      "standards_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_reason": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubStandardsControlAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubStandardsControlAssociation), &result)
	return &result
}
