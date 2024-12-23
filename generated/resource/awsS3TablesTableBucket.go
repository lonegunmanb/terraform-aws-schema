package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3TablesTableBucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "iceberg_unreferenced_file_removal": [
              "object",
              {
                "settings": [
                  "object",
                  {
                    "non_current_days": "number",
                    "unreferenced_days": "number"
                  }
                ],
                "status": "string"
              }
            ]
          }
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3TablesTableBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3TablesTableBucket), &result)
	return &result
}
