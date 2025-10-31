package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrpublicImages = `{
  "block": {
    "attributes": {
      "images": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "artifact_media_type": "string",
              "image_digest": "string",
              "image_manifest_media_type": "string",
              "image_pushed_at": "string",
              "image_size_in_bytes": "number",
              "image_tags": [
                "list",
                "string"
              ],
              "registry_id": "string",
              "repository_name": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_id": {
        "description": "AWS account ID associated with the public registry that contains the repository. If not specified, the default public registry is assumed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_name": {
        "description": "Name of the public repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "image_ids": {
        "block": {
          "attributes": {
            "image_digest": {
              "description": "Image digest.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_tag": {
              "description": "Image tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "List of image IDs to filter. Each image ID can use either a tag or digest.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description": "Provides details about AWS ECR Public Images in a public repository.",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrpublicImagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrpublicImages), &result)
	return &result
}
