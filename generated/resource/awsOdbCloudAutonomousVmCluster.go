package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudAutonomousVmCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "autonomous_data_storage_percentage": {
        "computed": true,
        "description": "The progress of the current operation on the Autonomous VM cluster, as a percentage.",
        "description_kind": "plain",
        "type": "number"
      },
      "autonomous_data_storage_size_in_tbs": {
        "description": "The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "available_autonomous_data_storage_size_in_tbs": {
        "computed": true,
        "description": "The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "available_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that you can create with the currently available storage.",
        "description_kind": "plain",
        "type": "number"
      },
      "available_cpus": {
        "computed": true,
        "description": "The number of CPU cores available for allocation to Autonomous Databases",
        "description_kind": "plain",
        "type": "number"
      },
      "cloud_exadata_infrastructure_id": {
        "description": "Exadata infrastructure id. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": "The compute model of the Autonomous VM cluster: ECPU or OCPU.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "The total number of CPU cores in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "cpu_core_count_per_node": {
        "description": "The number of CPU cores enabled per node in the Autonomous VM cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "cpu_percentage": {
        "computed": true,
        "description": "The percentage of total CPU cores currently in use in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the Autonomous VM cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_storage_size_in_gbs": {
        "computed": true,
        "description": "The total data storage allocated to the Autonomous VM cluster, in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description": "The total data storage allocated to the Autonomous VM cluster, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_servers": {
        "description": "The database servers in the Autonomous VM cluster. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "description": "The description of the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the Autonomous VM cluster. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain name of the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "exadata_storage_in_tbs_lowest_scaled_value": {
        "computed": true,
        "description": "The minimum value to which you can scale down the Exadata storage, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "hostname": {
        "computed": true,
        "description": "The hostname of the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_mtls_enabled_vm_cluster": {
        "computed": true,
        "description": "Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster. Changing this will force terraform to create new resource. ",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "license_model": {
        "computed": true,
        "description": "The license model for the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE . Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_acds_lowest_scaled_value": {
        "computed": true,
        "description": "The minimum value to which you can scale down the maximum number of Autonomous CDBs.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_per_oracle_compute_unit_in_gbs": {
        "description": "The amount of memory allocated per Oracle Compute Unit, in GB. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The total amount of memory allocated to the Autonomous VM cluster, in gigabytes(GB).",
        "description_kind": "plain",
        "type": "number"
      },
      "node_count": {
        "computed": true,
        "description": "The number of database server nodes in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "non_provisionable_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that can't be provisioned because of resource constraints.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor associated with this Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The URL for accessing the OCI console page for this Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "description": "The unique identifier of the ODB network associated with this Autonomous VM Cluster. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "odb_node_storage_size_in_gbs": {
        "computed": true,
        "description": " The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB)",
        "description_kind": "plain",
        "type": "number"
      },
      "percent_progress": {
        "computed": true,
        "description": "The progress of the current operation on the Autonomous VM cluster, as a percentage.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisionable_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_cpus": {
        "computed": true,
        "description": "The number of CPUs provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "reclaimable_cpus": {
        "computed": true,
        "description": "The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases.",
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_cpus": {
        "computed": true,
        "description": "The number of CPU cores reserved for system operations and redundancy.",
        "description_kind": "plain",
        "type": "number"
      },
      "scan_listener_port_non_tls": {
        "description": "The SCAN listener port for non-TLS (TCP) protocol. The default is 1521. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "scan_listener_port_tls": {
        "description": "The SCAN listener port for TLS (TCP) protocol. The default is 2484. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "shape": {
        "computed": true,
        "description": "The shape of the Exadata infrastructure for the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Autonomous VM cluster. Possible values include CREATING, AVAILABLE , UPDATING , DELETING , DELETED , FAILED ",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the current status of the Autonomous VM cluster.",
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
      "time_database_ssl_certificate_expires": {
        "computed": true,
        "description": "The expiration date and time of the database SSL certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "time_ords_certificate_expires": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description": "The time zone of the Autonomous VM cluster. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "total_container_databases": {
        "description": "The total number of Autonomous Container Databases that can be created with the allocated local storage. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "maintenance_window": {
        "block": {
          "attributes": {
            "days_of_week": {
              "description": "The days of the week when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ]
            },
            "hours_of_day": {
              "description": "The hours of the day when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "lead_time_in_weeks": {
              "description": "The lead time in weeks before the maintenance window.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "months": {
              "description": "The months when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ]
            },
            "preference": {
              "description": "The preference for the maintenance window scheduling.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weeks_of_month": {
              "description": "Indicates whether to skip release updates during maintenance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            }
          },
          "description": "The maintenance window of the Autonomous VM cluster.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbCloudAutonomousVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudAutonomousVmCluster), &result)
	return &result
}
