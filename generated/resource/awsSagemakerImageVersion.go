package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerImageVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "base_image": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "container_image": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "horovod": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ml_framework": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "processor": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "programming_lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_notes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vendor_guidance": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerImageVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerImageVersion), &result)
	return &result
}
