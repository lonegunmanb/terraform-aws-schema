package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsAlias = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKmsAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKmsAlias), &result)
	return &result
}
