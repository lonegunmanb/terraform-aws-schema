package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaRuntimeManagementConfig = `{
  "block": {
    "attributes": {
      "function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "qualifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_version_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_runtime_on": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaRuntimeManagementConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaRuntimeManagementConfig), &result)
	return &result
}
