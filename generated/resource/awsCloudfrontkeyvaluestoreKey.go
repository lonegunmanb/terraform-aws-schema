package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontkeyvaluestoreKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "description": "The key to put.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "key_value_store_arn": {
        "description": "The Amazon Resource Name (ARN) of the Key Value Store.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "total_size_in_bytes": {
        "computed": true,
        "description": "Total size of the Key Value Store in bytes.",
        "description_kind": "markdown",
        "type": "number"
      },
      "value": {
        "description": "The value to put.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudfrontkeyvaluestoreKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontkeyvaluestoreKey), &result)
	return &result
}
