package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchContributorManagedInsightRules = `{
  "block": {
    "attributes": {
      "managed_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_arn": "string",
              "rule_state": [
                "list",
                [
                  "object",
                  {
                    "rule_name": "string",
                    "state": "string"
                  }
                ]
              ],
              "template_name": "string"
            }
          ]
        ]
      },
      "resource_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchContributorManagedInsightRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchContributorManagedInsightRules), &result)
	return &result
}
