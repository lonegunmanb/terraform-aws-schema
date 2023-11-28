package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApprunnerDefaultAutoScalingConfigurationVersion = `{
  "block": {
    "attributes": {
      "auto_scaling_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApprunnerDefaultAutoScalingConfigurationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApprunnerDefaultAutoScalingConfigurationVersion), &result)
	return &result
}
