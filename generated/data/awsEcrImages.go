package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrImages = `{
  "block": {
    "attributes": {
      "image_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "image_digest": "string",
              "image_tag": "string"
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
        "description": "ID of the registry (AWS account ID)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_name": {
        "description": "Name of the repository",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrImagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrImages), &result)
	return &result
}
