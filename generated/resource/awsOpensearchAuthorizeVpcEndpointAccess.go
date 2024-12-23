package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchAuthorizeVpcEndpointAccess = `{
  "block": {
    "attributes": {
      "account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorized_principal": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal": "string",
              "principal_type": "string"
            }
          ]
        ]
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchAuthorizeVpcEndpointAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchAuthorizeVpcEndpointAccess), &result)
	return &result
}
