package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsXrayResourcePolicy = `{
  "block": {
    "attributes": {
      "bypass_policy_lockout_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_revision_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsXrayResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsXrayResourcePolicy), &result)
	return &result
}
