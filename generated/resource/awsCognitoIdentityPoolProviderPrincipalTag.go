package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoIdentityPoolProviderPrincipalTag = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_defaults": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoIdentityPoolProviderPrincipalTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoIdentityPoolProviderPrincipalTag), &result)
	return &result
}
