package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Export = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "body": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "export_version": {
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
      "include_extensions": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "output_type": {
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
      "specification": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2ExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Export), &result)
	return &result
}
