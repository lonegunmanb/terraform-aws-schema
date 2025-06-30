package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecretsmanagerSecretRotation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
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
      "rotate_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "rotation_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "rotation_lambda_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rotation_rules": {
        "block": {
          "attributes": {
            "automatically_after_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsSecretsmanagerSecretRotationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecretsmanagerSecretRotation), &result)
	return &result
}
