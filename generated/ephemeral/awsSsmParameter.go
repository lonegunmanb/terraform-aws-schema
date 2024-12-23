package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmParameter = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "with_decryption": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmParameterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmParameter), &result)
	return &result
}
