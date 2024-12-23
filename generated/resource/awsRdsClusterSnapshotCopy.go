package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsClusterSnapshotCopy = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "copy_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "db_cluster_snapshot_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "presigned_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "shared_accounts": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_cluster_snapshot_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "target_db_cluster_snapshot_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AwsRdsClusterSnapshotCopySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsClusterSnapshotCopy), &result)
	return &result
}
