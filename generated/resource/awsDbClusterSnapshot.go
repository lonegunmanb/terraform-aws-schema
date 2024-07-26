package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDbClusterSnapshot = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_snapshot_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_snapshot_identifier": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "source_db_cluster_snapshot_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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

func AwsDbClusterSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbClusterSnapshot), &result)
	return &result
}
