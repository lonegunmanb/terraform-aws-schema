package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSnsSmsPreferences = `{
  "block": {
    "attributes": {
      "default_sender_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_sms_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_status_iam_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_status_success_sampling_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monthly_spend_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "usage_report_s3_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSnsSmsPreferencesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSnsSmsPreferences), &result)
	return &result
}
