package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedpermissionsPolicyStore = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validation_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mode": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedpermissionsPolicyStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedpermissionsPolicyStore), &result)
	return &result
}
