package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBackupPlan = `{
  "block": {
    "attributes": {
      "arn": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "plan_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "completion_window": "number",
              "copy_action": [
                "set",
                [
                  "object",
                  {
                    "destination_vault_arn": "string",
                    "lifecycle": [
                      "list",
                      [
                        "object",
                        {
                          "cold_storage_after": "number",
                          "delete_after": "number",
                          "opt_in_to_archive_for_supported_resources": "bool"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "enable_continuous_backup": "bool",
              "lifecycle": [
                "list",
                [
                  "object",
                  {
                    "cold_storage_after": "number",
                    "delete_after": "number",
                    "opt_in_to_archive_for_supported_resources": "bool"
                  }
                ]
              ],
              "recovery_point_tags": [
                "map",
                "string"
              ],
              "rule_name": "string",
              "schedule": "string",
              "start_window": "number",
              "target_vault_name": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBackupPlan), &result)
	return &result
}
