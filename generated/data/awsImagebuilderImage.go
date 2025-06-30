package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderImage = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "build_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_recipe_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_recipe_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_scanning_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ecr_configuration": [
                "list",
                [
                  "object",
                  {
                    "container_tags": [
                      "set",
                      "string"
                    ],
                    "repository_name": "string"
                  }
                ]
              ],
              "image_scanning_enabled": "bool"
            }
          ]
        ]
      },
      "image_tests_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "image_tests_enabled": "bool",
              "timeout_minutes": "number"
            }
          ]
        ]
      },
      "infrastructure_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "output_resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "amis": [
                "set",
                [
                  "object",
                  {
                    "account_id": "string",
                    "description": "string",
                    "image": "string",
                    "name": "string",
                    "region": "string"
                  }
                ]
              ],
              "containers": [
                "set",
                [
                  "object",
                  {
                    "image_uris": [
                      "set",
                      "string"
                    ],
                    "region": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "platform": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsImagebuilderImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderImage), &result)
	return &result
}
