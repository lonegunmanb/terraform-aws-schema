package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaFunctionRecursionConfig = `{
  "block": {
    "attributes": {
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recursive_loop": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaFunctionRecursionConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaFunctionRecursionConfig), &result)
	return &result
}
