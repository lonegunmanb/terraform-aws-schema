package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKendraExperience = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "content_source_configuration": [
                "list",
                [
                  "object",
                  {
                    "data_source_ids": [
                      "set",
                      "string"
                    ],
                    "direct_put_content": "bool",
                    "faq_ids": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "user_identity_configuration": [
                "list",
                [
                  "object",
                  {
                    "identity_attribute_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "endpoint": "string",
              "endpoint_type": "string"
            }
          ]
        ]
      },
      "error_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "experience_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "index_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKendraExperienceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKendraExperience), &result)
	return &result
}
