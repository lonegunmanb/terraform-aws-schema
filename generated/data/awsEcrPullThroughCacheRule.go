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
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_registry_url": {
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
