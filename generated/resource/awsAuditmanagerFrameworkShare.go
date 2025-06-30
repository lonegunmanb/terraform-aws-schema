package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAuditmanagerFrameworkShare = `{
  "block": {
    "attributes": {
      "comment": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "framework_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAuditmanagerFrameworkShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAuditmanagerFrameworkShare), &result)
	return &result
}
