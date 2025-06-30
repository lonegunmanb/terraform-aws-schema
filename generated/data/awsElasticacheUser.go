package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheUser = `{
  "block": {
    "attributes": {
      "access_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
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
      "no_password_required": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "passwords": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "set",
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
      "user_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "authentication_mode": {
        "block": {
          "attributes": {
            "password_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheUser), &result)
	return &result
}
