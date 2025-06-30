package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrPullThroughCacheRule = `{
  "block": {
    "attributes": {
      "credential_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ecr_repository_prefix": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
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
      },
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_registry_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_repository_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrPullThroughCacheRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrPullThroughCacheRule), &result)
	return &result
}
