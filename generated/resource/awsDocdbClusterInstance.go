package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDocdbClusterInstance = `{
  "block": {
    "attributes": {
      "apply_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_cert_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dbi_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_performance_insights": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "optional": true,
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
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "promotion_tier": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "writer": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsDocdbClusterInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDocdbClusterInstance), &result)
	return &result
}
