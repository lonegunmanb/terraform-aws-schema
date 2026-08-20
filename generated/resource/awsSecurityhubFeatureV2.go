package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubFeatureV2 = `{
  "block": {
    "attributes": {
      "feature_name": {
        "description": "The name of the opt-in feature to enable. Valid values: NETWORK_SCANNING.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "feature_status": {
        "description": "The current enablement status of the feature. Valid values: ENABLED, DISABLED.",
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

func AwsSecurityhubFeatureV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubFeatureV2), &result)
	return &result
}
