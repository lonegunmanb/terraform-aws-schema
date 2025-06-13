package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessSecurityPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Description of the policy. Typically used to store information about the permissions defined in the policy.",
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
