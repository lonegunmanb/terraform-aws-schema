package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchContributorInsightRule = `{
  "block": {
    "attributes": {
      "resource_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_definition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchContributorInsightRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchContributorInsightRule), &result)
	return &result
}
