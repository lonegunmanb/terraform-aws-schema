package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsArcregionswitchRoute53HealthChecks = `{
  "block": {
    "attributes": {
      "health_checks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "health_check_id": "string",
              "hosted_zone_id": "string",
              "record_name": "string",
              "region": "string",
              "status": "string"
            }
          ]
        ]
      },
      "plan_arn": {
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

func AwsArcregionswitchRoute53HealthChecksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsArcregionswitchRoute53HealthChecks), &result)
	return &result
}
