package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsInspector2OrganizationConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_account_limit_reached": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "auto_enable": {
        "block": {
          "attributes": {
            "code_repository": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ec2": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "ecr": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "lambda": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "lambda_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsInspector2OrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsInspector2OrganizationConfiguration), &result)
	return &result
}
