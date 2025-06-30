package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2ApiKey = `{
  "block": {
    "attributes": {
      "api_key": {
        "computed": true,
        "description": "The API key value. This is sensitive and not included in responses.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "token_domains": {
        "description": "The domains that you want to be able to use the API key with, for example example.com. Maximum of 5 domains.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Provides a WAFv2 API Key resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWafv2ApiKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2ApiKey), &result)
	return &result
}
