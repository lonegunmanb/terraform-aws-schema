package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubEnabledStandards = `{
  "block": {
    "attributes": {
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "standards_subscription_arns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "standards_subscriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "standards_arn": "string",
              "standards_controls_updatable": "string",
              "standards_inputs": [
                "map",
                "string"
              ],
              "standards_status": "string",
              "standards_status_reason": [
                "list",
                [
                  "object",
                  {
                    "status_reason_code": "string"
                  }
                ]
              ],
              "standards_subscription_arn": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubEnabledStandardsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubEnabledStandards), &result)
	return &result
}
