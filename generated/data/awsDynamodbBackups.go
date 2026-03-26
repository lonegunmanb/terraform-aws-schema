package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDynamodbBackups = `{
  "block": {
    "attributes": {
      "backup_summaries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_arn": "string",
              "backup_creation_date_time": "string",
              "backup_expiry_date_time": "string",
              "backup_name": "string",
              "backup_size_bytes": "number",
              "backup_status": "string",
              "backup_type": "string",
              "table_arn": "string",
              "table_id": "string",
              "table_name": "string"
            }
          ]
        ]
      },
      "backup_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_range_lower_bound": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_range_upper_bound": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDynamodbBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDynamodbBackups), &result)
	return &result
}
