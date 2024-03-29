package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsResourceexplorer2Search = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "view_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "resource_count": {
        "block": {
          "attributes": {
            "complete": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "total_resources": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "resources": {
        "block": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "last_reported_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "owning_account_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "resource_property": {
              "block": {
                "attributes": {
                  "data": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "last_reported_at": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
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

func AwsResourceexplorer2SearchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsResourceexplorer2Search), &result)
	return &result
}
