package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsCluster = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allow_major_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "apply_immediately": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "backtrack_window": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_certificate_valid_till": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_identifier_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_members": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cluster_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_scalability_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "database_insights_mode": {
        "computed": true,
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
      "db_cluster_instance_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_parameter_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_system_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_automated_backups": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_iam_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_global_write_forwarding": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_http_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_local_write_forwarding": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled_cloudwatch_logs_exports": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_lifecycle_support": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version_actual": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "final_snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_cluster_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iam_database_authentication_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "iops": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manage_master_user_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "master_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "master_user_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_id": "string",
              "secret_arn": "string",
              "secret_status": "string"
            }
          ]
        ]
      },
      "master_user_secret_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_username": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitoring_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_insights_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reader_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_source_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_final_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_type": {
        "computed": true,
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
      "restore_to_point_in_time": {
        "block": {
          "attributes": {
            "restore_to_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "restore_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_cluster_identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_cluster_resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_latest_restorable_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "s3_import": {
        "block": {
          "attributes": {
            "bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bucket_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ingestion_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_engine": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_engine_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scaling_configuration": {
        "block": {
          "attributes": {
            "auto_pause": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "seconds_before_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "seconds_until_auto_pause": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "serverlessv2_scaling_configuration": {
        "block": {
          "attributes": {
            "max_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "seconds_until_auto_pause": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
  "version": 1
}`

func AwsRdsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsCluster), &result)
	return &result
}
