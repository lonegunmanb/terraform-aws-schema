package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAthenaNamedQuery = `{
  "block": {
    "attributes": {
      "database": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
        "required": true,
        "type": "string"
      },
      "querystring": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workgroup": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAthenaNamedQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAthenaNamedQuery), &result)
	return &result
}
