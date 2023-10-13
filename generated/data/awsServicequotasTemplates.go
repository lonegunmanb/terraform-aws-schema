package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicequotasTemplates = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "templates": {
        "block": {
          "attributes": {
            "global_quota": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "quota_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "quota_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "unit": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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

func AwsServicequotasTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicequotasTemplates), &result)
	return &result
}
