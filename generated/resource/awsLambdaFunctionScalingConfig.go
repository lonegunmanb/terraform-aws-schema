package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaFunctionScalingConfig = `{
  "block": {
    "attributes": {
      "function_arn": {
        "computed": true,
        "description": "ARN of the Lambda function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description": "Name or ARN of the Lambda function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_state": {
        "computed": true,
        "description": "State of the function after applying the scaling configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "qualifier": {
        "description": "Qualifier for the scaling configuration. Valid values: $LATEST.PUBLISHED or a numeric version number.",
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
      "function_scaling_config": {
        "block": {
          "attributes": {
            "max_execution_environments": {
              "computed": true,
              "description": "Maximum number of execution environments that can be provisioned for the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_execution_environments": {
              "computed": true,
              "description": "Minimum number of execution environments to maintain for the function.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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

func AwsLambdaFunctionScalingConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaFunctionScalingConfig), &result)
	return &result
}
