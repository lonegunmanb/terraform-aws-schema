package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaInvocation = `{
  "block": {
    "attributes": {
      "client_context": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "executed_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_error": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_result": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payload": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "qualifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "result": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaInvocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaInvocation), &result)
	return &result
}
