package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbAutonomousDatabase = `{
  "block": {
    "attributes": {
      "actual_used_data_storage_size_in_tbs": {
        "computed": true,
        "description": "Actual amount of data storage currently in use, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "admin_password": {
        "description": "Password for the ADMIN user. This value is stored in Terraform state. Use admin_password_wo with Terraform 1.11 or later to avoid storing the password in state.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "admin_password_wo": {
        "description": "Password for the ADMIN user. This write-only value is never stored in Terraform state.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string",
        "write_only": true
      },
      "admin_password_wo_version": {
        "description": "Arbitrary version used to trigger an ADMIN password update.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocated_storage_size_in_tbs": {
        "computed": true,
        "description": "Amount of storage currently allocated, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "allowlisted_ips": {
        "computed": true,
        "description": "IP addresses allowed to access the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_refresh_frequency_in_seconds": {
        "computed": true,
        "description": "Frequency at which a refreshable clone is automatically refreshed, in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_refresh_point_lag_in_seconds": {
        "computed": true,
        "description": "Time lag between a refreshable clone and its source, in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "autonomous_maintenance_schedule_type": {
        "computed": true,
        "description": "Maintenance schedule type for the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "Availability Zone where the Autonomous Database is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "Availability Zone ID where the Autonomous Database is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "available_upgrade_versions": {
        "computed": true,
        "description": "Oracle Database versions to which the Autonomous Database can be upgraded.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backup_retention_period_in_days": {
        "computed": true,
        "description": "Retention period for automatic backups, in days.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "byol_compute_count_limit": {
        "computed": true,
        "description": "Maximum compute capacity under the bring-your-own-license model.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "character_set": {
        "computed": true,
        "description": "Character set of the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_count": {
        "computed": true,
        "description": "Compute capacity in ECPUs or OCPUs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description": "Compute model of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "Number of CPU cores allocated to the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "Date and time when the Autonomous Database was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_storage_size_in_gbs": {
        "computed": true,
        "description": "Data volume size in GB.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description": "Data volume size in TB. Configured values must be whole numbers; computed values may be fractional when storage is configured in GB.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "database_edition": {
        "computed": true,
        "description": "Oracle Database edition.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_type": {
        "computed": true,
        "description": "Type of Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "Name of the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_version": {
        "computed": true,
        "description": "Oracle Database software version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_workload": {
        "computed": true,
        "description": "Intended database workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "User-friendly name for the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key_provider": {
        "computed": true,
        "description": "Encryption key provider. Configurable values are ORACLE_MANAGED and AWS_KMS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_auto_scaling_enabled": {
        "computed": true,
        "description": "Whether automatic compute scaling is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_auto_scaling_for_storage_enabled": {
        "computed": true,
        "description": "Whether automatic storage scaling is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_backup_retention_locked": {
        "computed": true,
        "description": "Whether the backup retention period is locked.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_local_data_guard_enabled": {
        "computed": true,
        "description": "Whether local Oracle Data Guard is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_mtls_connection_required": {
        "computed": true,
        "description": "Whether mutual TLS authentication is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_refreshable_clone": {
        "computed": true,
        "description": "Whether the Autonomous Database is a refreshable clone.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_id": {
        "computed": true,
        "description": "ARN of the AWS KMS key used to encrypt the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "Oracle license model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_adg_auto_failover_max_data_loss_limit": {
        "computed": true,
        "description": "Maximum data-loss limit for automatic local Data Guard failover, in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ncharacter_set": {
        "computed": true,
        "description": "National character set of the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "Name of the OCI resource anchor.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "URL for the Autonomous Database in the OCI console.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "Oracle Cloud Identifier of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "ARN of the associated ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "ID of the associated ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "open_mode": {
        "computed": true,
        "description": "Open mode of the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "Progress of the current operation, as a percentage.",
        "description_kind": "plain",
        "type": "number"
      },
      "permission_level": {
        "computed": true,
        "description": "Permission level of the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_endpoint": {
        "computed": true,
        "description": "Private endpoint of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint_ip": {
        "computed": true,
        "description": "Private endpoint IP address.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_endpoint_label": {
        "computed": true,
        "description": "Private endpoint label.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "refreshable_mode": {
        "computed": true,
        "description": "Refresh mode of a refreshable clone.",
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
      "resource_pool_leader_id": {
        "computed": true,
        "description": "ID of the resource-pool leader Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_console_url": {
        "computed": true,
        "description": "URL for the Oracle service console.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "description": "Source from which to create the Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_id": {
        "computed": true,
        "description": "ID of the source used to create the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "sql_web_developer_url": {
        "computed": true,
        "description": "URL for Oracle SQL Developer Web.",
        "description_kind": "plain",
        "type": "string"
      },
      "standby_allowlisted_ips": {
        "computed": true,
        "description": "IP addresses allowed to access the standby Autonomous Database.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "standby_allowlisted_ips_source": {
        "computed": true,
        "description": "Source of the standby allowlisted IP addresses.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the current status.",
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
      "time_of_auto_refresh_start": {
        "computed": true,
        "description": "Date and time when automatic refresh begins.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "admin_password_source": {
        "block": {
          "block_types": {
            "customer_managed_aws_secret": {
              "block": {
                "attributes": {
                  "external_id_type": {
                    "description": "Type of OCI identifier supplied as the external ID when OCI assumes the IAM role.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "iam_role_arn": {
                    "description": "ARN of the customer-managed IAM role OCI assumes to retrieve the secret.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_arn": {
                    "description": "ARN of the AWS Secrets Manager secret that contains the ADMIN password.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Source of the ADMIN password. Conflicts with admin_password and admin_password_wo.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "customer_contacts_to_send_to_oci": {
        "block": {
          "attributes": {
            "email": {
              "description": "Email address of the customer contact.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Customer contacts that receive operational notifications from OCI.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "db_tools_details": {
        "block": {
          "attributes": {
            "compute_count": {
              "computed": true,
              "description": "Compute capacity allocated to the database tool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "is_enabled": {
              "computed": true,
              "description": "Whether the database tool is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_idle_time_in_minutes": {
              "computed": true,
              "description": "Maximum idle time before the database tool is shut down, in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "Name of the database tool.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Database management tools enabled for the Autonomous Database.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "long_term_backup_schedule": {
        "block": {
          "attributes": {
            "is_disabled": {
              "computed": true,
              "description": "Whether the long-term backup schedule is disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "repeat_cadence": {
              "computed": true,
              "description": "Cadence at which long-term backups are taken.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retention_period_in_days": {
              "computed": true,
              "description": "Retention period for long-term backups, in days.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_of_backup": {
              "computed": true,
              "description": "Date and time at which the long-term backup is taken.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Long-term backup schedule for the Autonomous Database.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "resource_pool_summary": {
        "block": {
          "attributes": {
            "available_compute_capacity": {
              "computed": true,
              "description": "Available compute capacity in the resource pool.",
              "description_kind": "plain",
              "type": "number"
            },
            "available_storage_capacity_in_tbs": {
              "computed": true,
              "description": "Available storage capacity in the resource pool, in TB.",
              "description_kind": "plain",
              "type": "number"
            },
            "is_disabled": {
              "computed": true,
              "description": "Whether the resource pool is disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pool_size": {
              "computed": true,
              "description": "Number of Autonomous Databases the resource pool can contain.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pool_storage_size_in_tbs": {
              "computed": true,
              "description": "Total storage size of the resource pool, in TB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total_compute_capacity": {
              "computed": true,
              "description": "Total compute capacity of the resource pool.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description": "Resource pool configuration for the Autonomous Database.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "scheduled_operations": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description": "Day of the week.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scheduled_start_time": {
              "computed": true,
              "description": "Scheduled start time in UTC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduled_stop_time": {
              "computed": true,
              "description": "Scheduled stop time in UTC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Scheduled start and stop times for the Autonomous Database.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "source_configuration": {
        "block": {
          "block_types": {
            "clone_to_refreshable": {
              "block": {
                "attributes": {
                  "auto_refresh_frequency_in_seconds": {
                    "description": "Frequency at which the refreshable clone is automatically refreshed, in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "auto_refresh_point_lag_in_seconds": {
                    "description": "Time lag between the refreshable clone and its source, in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "clone_type": {
                    "description": "Type of clone to create.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "open_mode": {
                    "description": "Open mode of the refreshable clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "refreshable_mode": {
                    "description": "Refresh mode of the clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_autonomous_database_id": {
                    "description": "ID of the source Autonomous Database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time_of_auto_refresh_start": {
                    "description": "Date and time when automatic refresh starts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cross_region_data_guard": {
              "block": {
                "attributes": {
                  "source_autonomous_database_arn": {
                    "description": "ARN of the source Autonomous Database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cross_region_disaster_recovery": {
              "block": {
                "attributes": {
                  "is_replicate_automatic_backups": {
                    "description": "Whether automatic backups are replicated to the disaster recovery database.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "remote_disaster_recovery_type": {
                    "description": "Type of remote disaster recovery.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_autonomous_database_arn": {
                    "description": "ARN of the source Autonomous Database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "database_clone": {
              "block": {
                "attributes": {
                  "clone_type": {
                    "description": "Type of clone to create.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_autonomous_database_id": {
                    "description": "ID of the source Autonomous Database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "point_in_time_restore": {
              "block": {
                "attributes": {
                  "clone_table_space_list": {
                    "description": "Tablespace IDs to clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "clone_type": {
                    "description": "Type of clone to create.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_autonomous_database_id": {
                    "description": "ID of the source Autonomous Database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timestamp": {
                    "description": "Date and time to which the database is restored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "use_latest_available_backup_timestamp": {
                    "description": "Whether to use the latest available backup timestamp.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "restore_from_backup": {
              "block": {
                "attributes": {
                  "autonomous_database_backup_id": {
                    "description": "ID of the Autonomous Database backup.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "clone_table_space_list": {
                    "description": "Tablespace IDs to clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "clone_type": {
                    "description": "Type of clone to create from the backup.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Source-specific configuration used during creation. Exactly one nested source block must be configured.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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
      },
      "transportable_tablespace": {
        "block": {
          "attributes": {
            "tts_bundle_url": {
              "description": "URL of the transportable tablespace bundle.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Transportable tablespace configuration used during creation.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbAutonomousDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbAutonomousDatabase), &result)
	return &result
}
