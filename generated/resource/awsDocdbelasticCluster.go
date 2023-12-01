package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDocdbelasticCluster = `{
  "block": {
    "attributes": {
      "admin_user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "admin_user_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
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
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "shard_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "shard_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDocdbelasticClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDocdbelasticCluster), &result)
	return &result
}
