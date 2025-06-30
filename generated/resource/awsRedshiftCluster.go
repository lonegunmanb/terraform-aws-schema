package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftCluster = `{
  "block": {
    "attributes": {
      "allow_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "apply_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "aqua_configuration_status": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "automated_snapshot_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_relocation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_namespace_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "node_role": "string",
              "private_ip_address": "string",
              "public_ip_address": "string"
            }
          ]
        ]
      },
      "cluster_parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_revision_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_subnet_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_iam_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "final_snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_roles": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
        "optional": true,
        "type": "string"
      },
      "maintenance_track_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manage_master_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "manual_snapshot_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "master_password_secret_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_password_secret_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_password_wo": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string",
        "write_only": true
      },
      "master_password_wo_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_username": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "node_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "owner_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_final_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_cluster_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
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

func AwsRedshiftClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftCluster), &result)
	return &result
}
