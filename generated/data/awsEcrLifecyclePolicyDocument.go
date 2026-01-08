package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrLifecyclePolicyDocument = `{
  "block": {
    "attributes": {
      "json": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "target_storage_class": {
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
              "nesting_mode": "list"
            },
            "selection": {
              "block": {
                "attributes": {
                  "count_number": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "count_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "count_unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_class": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_pattern_list": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tag_prefix_list": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tag_status": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcrLifecyclePolicyDocumentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrLifecyclePolicyDocument), &result)
	return &result
}
