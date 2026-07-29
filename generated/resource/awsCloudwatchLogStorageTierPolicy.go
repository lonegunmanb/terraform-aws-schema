package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchLogStorageTierPolicy = `{
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
      },
      "storage_tier": {
        "description": "The storage tier to set for the account. Valid values are ` + "`" + `STANDARD` + "`" + ` or ` + "`" + `INTELLIGENT_TIERING` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Manages a CloudWatch Logs account-level storage tier policy. When set to ` + "`" + `INTELLIGENT_TIERING` + "`" + `, CloudWatch Logs automatically moves log data to the most cost-effective storage tier based on access frequency.",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchLogStorageTierPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchLogStorageTierPolicy), &result)
	return &result
}
