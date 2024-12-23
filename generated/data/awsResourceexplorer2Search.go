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
      "resource_count": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "complete": "bool",
              "total_resources": "number"
            }
          ]
        ]
      },
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "last_reported_at": "string",
              "owning_account_id": "string",
              "properties": [
                "list",
                [
                  "object",
                  {
                    "data": "string",
                    "last_reported_at": "string",
                    "name": "string"
                  }
                ]
              ],
              "region": "string",
              "resource_type": "string",
              "service": "string"
            }
          ]
        ]
      },
      "view_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
