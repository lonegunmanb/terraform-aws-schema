package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPollyVoices = `{
  "block": {
    "attributes": {
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "include_additional_language_codes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "language_code": {
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
      }
    },
    "block_types": {
      "voices": {
        "block": {
          "attributes": {
            "additional_language_codes": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "gender": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "language_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "language_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "supported_engines": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
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

func AwsPollyVoicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPollyVoices), &result)
	return &result
}
