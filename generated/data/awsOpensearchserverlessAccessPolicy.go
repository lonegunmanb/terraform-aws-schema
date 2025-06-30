package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessAccessPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the policy. Typically used to store information about the permissions defined in the policy.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "JSON policy document to use as the content for the new policy.",
        "description_kind": "plain",
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
        "description": "Type of access policy. Must be ` + "`" + `data` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessAccessPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessAccessPolicy), &result)
	return &result
}
