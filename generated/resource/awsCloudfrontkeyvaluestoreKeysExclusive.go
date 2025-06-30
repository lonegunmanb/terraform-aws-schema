package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontkeyvaluestoreKeysExclusive = `{
  "block": {
    "attributes": {
      "key_value_store_arn": {
        "description": "The Amazon Resource Name (ARN) of the Key Value Store.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "max_batch_size": {
        "computed": true,
        "description": "Maximum resource key values pairs that you wills update in a single API request. AWS has a default quota of 50 keys or a 3 MB payload, whichever is reached first",
        "description_kind": "markdown",
        "optional": true,
        "type": "number"
      },
      "total_size_in_bytes": {
        "computed": true,
        "description": "Total size of the Key Value Store in bytes.",
        "description_kind": "markdown",
        "type": "number"
      }
    },
    "block_types": {
      "resource_key_value_pair": {
        "block": {
          "attributes": {
            "key": {
              "description": "The key to put.",
              "description_kind": "markdown",
              "required": true,
              "type": "string"
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudfrontkeyvaluestoreKeysExclusiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontkeyvaluestoreKeysExclusive), &result)
	return &result
}
