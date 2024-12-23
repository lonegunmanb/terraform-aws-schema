package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftSnapshotCopy = `{
  "block": {
    "attributes": {
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "manual_snapshot_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot_copy_grant_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftSnapshotCopySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftSnapshotCopy), &result)
	return &result
}
