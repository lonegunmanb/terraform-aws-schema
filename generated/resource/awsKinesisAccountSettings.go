package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKinesisAccountSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
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
    "block_types": {
      "minimum_throughput_billing_commitment": {
        "block": {
          "attributes": {
            "earliest_allowed_end_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ended_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "started_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status_actual": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKinesisAccountSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKinesisAccountSettings), &result)
	return &result
}
