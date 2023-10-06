package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicequotasTemplate = `{
  "block": {
    "attributes": {
      "global_quota": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "quota_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quota_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_code": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicequotasTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicequotasTemplate), &result)
	return &result
}
