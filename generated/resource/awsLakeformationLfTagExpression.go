package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLakeformationLfTagExpression = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The ID of the Data Catalog.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A description of the LF-Tag Expression.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the LF-Tag Expression.",
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
    "block_types": {
      "expression": {
        "block": {
          "attributes": {
            "tag_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tag_values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description": "Manages an AWS Lake Formation Tag Expression.",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLakeformationLfTagExpressionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLakeformationLfTagExpression), &result)
	return &result
}
