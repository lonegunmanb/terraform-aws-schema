package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkmailGroup = `{
  "block": {
    "attributes": {
      "disabled_date": {
        "computed": true,
        "description": "Timestamp when the group was disabled from WorkMail use.",
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
        "description": "Primary email address used to register the group with WorkMail.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled_date": {
        "computed": true,
        "description": "Timestamp when the group was enabled for WorkMail use.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description": "Identifier of the group.",
        "description_kind": "plain",
        "type": "string"
      },
      "hidden_from_global_address_list": {
        "computed": true,
        "description": "Whether to hide the group from the global address list.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "Name of the group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_id": {
        "description": "Identifier of the WorkMail organization where the group is managed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current WorkMail state of the group.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWorkmailGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkmailGroup), &result)
	return &result
}
