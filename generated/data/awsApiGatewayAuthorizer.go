package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayAuthorizer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_credentials": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorizer_result_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "authorizer_uri": {
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
      "identity_source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_validation_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rest_api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApiGatewayAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayAuthorizer), &result)
	return &result
}
