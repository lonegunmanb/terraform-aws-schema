package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrImage = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_digest": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_pushed_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "image_size_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "image_tag": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "image_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrImage), &result)
	return &result
}
