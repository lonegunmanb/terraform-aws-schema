package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreApiKeyCredentialProvider = `{
  "block": {
    "attributes": {
      "api_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "api_key_secret_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "secret_arn": "string"
            }
          ]
        ]
      },
      "api_key_wo": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string",
        "write_only": true
      },
      "api_key_wo_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "credential_provider_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockagentcoreApiKeyCredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreApiKeyCredentialProvider), &result)
	return &result
}
