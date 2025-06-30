package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicequotasTemplates = `{
  "block": {
    "attributes": {
      "aws_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "global_quota": "bool",
              "quota_code": "string",
              "quota_name": "string",
              "region": "string",
              "service_code": "string",
              "service_name": "string",
              "unit": "string",
              "value": "number"
            }
          ]
        ]
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
