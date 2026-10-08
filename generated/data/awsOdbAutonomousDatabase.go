package data

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
      "admin_password_source": {
        "computed": true,
        "description": "Source of the ADMIN password.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "customer_managed_aws_secret": [
                "list",
                [
                  "object",
                  {
                    "external_id_type": "string",
                    "iam_role_arn": "string",
                    "secret_arn": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
        "description": "Automatic refresh frequency, in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "auto_refresh_point_lag_in_seconds": {
        "computed": true,
        "description": "Refresh lag from the source, in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "autonomous_maintenance_schedule_type": {
        "computed": true,
        "description": "Maintenance schedule type.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "Availability Zone of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "Availability Zone ID of the Autonomous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "available_upgrade_versions": {
        "computed": true,
        "description": "Oracle Database versions available for upgrade.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backup_retention_period_in_days": {
        "computed": true,
        "description": "Automatic backup retention period, in days.",
        "description_kind": "plain",
        "type": "number"
      },
      "byol_compute_count_limit": {
        "computed": true,
        "description": "Maximum BYOL compute capacity.",
        "description_kind": "plain",
        "type": "number"
      },
      "character_set": {
        "computed": true,
        "description": "Database character set.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_count": {
        "computed": true,
        "description": "Compute capacity in ECPUs or OCPUs.",
        "description_kind": "plain",
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description": "Compute model.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "Allocated CPU core count.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "Creation date and time.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_contacts_to_send_to_oci": {
        "computed": true,
        "description": "Customer contacts that receive operational notifications from OCI.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "email": "string"
            }
          ]
        ]
      },
      "data_storage_size_in_gbs": {
        "computed": true,
        "description": "Data volume size in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description": "Data volume size in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "database_edition": {
        "computed": true,
        "description": "Oracle Database edition.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_type": {
        "computed": true,
        "description": "Autonomous Database type.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "Database name.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_tools_details": {
        "computed": true,
        "description": "Database management tools enabled for the Autonomous Database.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "compute_count": "number",
              "is_enabled": "bool",
              "max_idle_time_in_minutes": "number",
              "name": "string"
            }
          ]
        ]
      },
      "db_version": {
        "computed": true,
        "description": "Oracle Database software version.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_workload": {
        "computed": true,
        "description": "Database workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "User-friendly database name.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_provider": {
        "computed": true,
        "description": "Encryption key provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Unique identifier of the Autonomous Database.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_auto_scaling_enabled": {
        "computed": true,
        "description": "Whether automatic compute scaling is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_auto_scaling_for_storage_enabled": {
        "computed": true,
        "description": "Whether automatic storage scaling is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_backup_retention_locked": {
        "computed": true,
        "description": "Whether backup retention is locked.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_local_data_guard_enabled": {
        "computed": true,
        "description": "Whether local Data Guard is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_mtls_connection_required": {
        "computed": true,
        "description": "Whether mTLS is required.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_refreshable_clone": {
        "computed": true,
        "description": "Whether the database is a refreshable clone.",
        "description_kind": "plain",
        "type": "bool"
      },
      "kms_key_id": {
        "computed": true,
        "description": "ARN of the AWS KMS encryption key.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "Oracle license model.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_adg_auto_failover_max_data_loss_limit": {
        "computed": true,
        "description": "Maximum automatic-failover data-loss limit, in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "long_term_backup_schedule": {
        "computed": true,
        "description": "Long-term backup schedule.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_disabled": "bool",
              "repeat_cadence": "string",
              "retention_period_in_days": "number",
              "time_of_backup": "string"
            }
          ]
        ]
      },
      "ncharacter_set": {
        "computed": true,
        "description": "National character set.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "OCI resource anchor name.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "OCI console URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "Oracle Cloud Identifier.",
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
        "type": "string"
      },
      "open_mode": {
        "computed": true,
        "description": "Database open mode.",
        "description_kind": "plain",
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "Progress of the current operation.",
        "description_kind": "plain",
        "type": "number"
      },
      "permission_level": {
        "computed": true,
        "description": "Database permission level.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint": {
        "computed": true,
        "description": "Private endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint_ip": {
        "computed": true,
        "description": "Private endpoint IP address.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint_label": {
        "computed": true,
        "description": "Private endpoint label.",
        "description_kind": "plain",
        "type": "string"
      },
      "refreshable_mode": {
        "computed": true,
        "description": "Refreshable clone mode.",
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
      "resource_pool_leader_id": {
        "computed": true,
        "description": "Resource-pool leader database ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_pool_summary": {
        "computed": true,
        "description": "Resource pool configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_compute_capacity": "number",
              "available_storage_capacity_in_tbs": "number",
              "is_disabled": "bool",
              "pool_size": "number",
              "pool_storage_size_in_tbs": "number",
              "total_compute_capacity": "number"
            }
          ]
        ]
      },
      "scheduled_operations": {
        "computed": true,
        "description": "Scheduled database start and stop times.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "day_of_week": "string",
              "scheduled_start_time": "string",
              "scheduled_stop_time": "string"
            }
          ]
        ]
      },
      "service_console_url": {
        "computed": true,
        "description": "Oracle service console URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_id": {
        "computed": true,
        "description": "Source database or backup ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "sql_web_developer_url": {
        "computed": true,
        "description": "Oracle SQL Developer Web URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "standby_allowlisted_ips": {
        "computed": true,
        "description": "IP addresses allowed to access the standby database.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "standby_allowlisted_ips_source": {
        "computed": true,
        "description": "Source of standby allowlisted IPs.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current database status.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional status information.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "time_of_auto_refresh_start": {
        "computed": true,
        "description": "Automatic refresh start time.",
        "description_kind": "plain",
        "type": "string"
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
