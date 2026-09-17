package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaResourcePolicy = `{
  "block": {
    "attributes": {
      "policy": {
        "description": "JSON-formatted resource-based policy document to attach to the Lambda resource. This replaces the entire policy, including any statements added with aws_lambda_permission.",
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
      },
      "resource_arn": {
        "description": "ARN of the Lambda function, version, or alias to attach the resource-based policy to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "Unique identifier for the current revision of the policy. Changes on every update, since PutResourcePolicy always issues a new revision.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaResourcePolicy), &result)
	return &result
}
