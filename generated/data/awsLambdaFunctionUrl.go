package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaFunctionUrl = `{
  "block": {
    "attributes": {
      "authorization_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cors": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_credentials": "bool",
              "allow_headers": [
                "list",
                "string"
              ],
              "allow_methods": [
                "list",
                "string"
              ],
              "allow_origins": [
                "list",
                "string"
              ],
              "expose_headers": [
                "list",
                "string"
              ],
              "max_age": "number"
            }
          ]
        ]
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
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
      "function_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "invoke_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
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
      "url_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaFunctionUrlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaFunctionUrl), &result)
	return &result
}
