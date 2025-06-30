package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessLifecyclePolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Description of the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "description": "JSON policy document to use as the content for the new policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_version": {
        "computed": true,
        "description": "Version of the policy.",
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
      "type": {
        "description": "Type of lifecycle policy. Must be ` + "`" + `retention` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessLifecyclePolicy), &result)
	return &result
}
