package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafregionalByteMatchSet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "byte_match_tuple": {
        "block": {
          "attributes": {
            "positional_constraint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "text_transformation": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "field_to_match": {
              "block": {
                "attributes": {
                  "data": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "byte_match_tuples": {
        "block": {
          "attributes": {
            "positional_constraint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "text_transformation": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "field_to_match": {
              "block": {
                "attributes": {
                  "data": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
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

func AwsWafregionalByteMatchSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafregionalByteMatchSet), &result)
	return &result
}
