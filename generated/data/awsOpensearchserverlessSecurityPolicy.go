package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessSecurityPolicy = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description": "The date the security policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the security policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "The date the security policy was last modified.",
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
        "description": "The JSON policy document without any whitespaces.",
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
        "description": "Type of security policy. One of ` + "`" + `encryption` + "`" + ` or ` + "`" + `network` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessSecurityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessSecurityPolicy), &result)
	return &result
}
