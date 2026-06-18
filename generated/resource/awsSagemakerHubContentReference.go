package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerHubContentReference = `{
  "block": {
    "attributes": {
      "hub_arn": {
        "computed": true,
        "description": "ARN of the private SageMaker Hub that contains the content reference.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_arn": {
        "computed": true,
        "description": "ARN of the hub content reference (without version suffix). The min_version is stripped off from the end of this ARN to make it usable to list tags.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_name": {
        "description": "Name of the hub content reference.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hub_content_status": {
        "computed": true,
        "description": "Status of the hub content reference. Valid values include ` + "`" + `Available` + "`" + `, ` + "`" + `Importing` + "`" + `, ` + "`" + `Deleting` + "`" + `, ` + "`" + `ImportFailed` + "`" + `, ` + "`" + `DeleteFailed` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_version": {
        "computed": true,
        "description": "Version of the hub content reference.",
        "description_kind": "plain",
        "type": "string"
      },
      "hub_name": {
        "description": "Name of the private SageMaker Hub to add the content reference to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "min_version": {
        "description": "Minimum version of the hub content to reference. Use \"1.0.0\" to support all versions. Changing this value to an empty string forces replacement of the resource.",
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
      "sagemaker_public_hub_content_arn": {
        "description": "ARN of the public SageMaker JumpStart hub content to reference. The ARN must not include a version suffix.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerHubContentReferenceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerHubContentReference), &result)
	return &result
}
