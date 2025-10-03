package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudVmCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "description": "The unique identifier of the Exadata infrastructure for this VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the Grid Infrastructure (GI) cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": "The compute model used when the instance is created or cloned — either ECPU or OCPU. ECPU is a virtualized compute unit; OCPU is a physical processor core with hyper-threading.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "description": "The number of CPU cores to enable on the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the VM cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_storage_size_in_tbs": {
        "description": "The size of the data disk group, in terabytes (TBs), to allocate for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GBs), to allocate for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_servers": {
        "description": "The list of database servers for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "disk_redundancy": {
        "computed": true,
        "description": "The type of redundancy for the VM cluster: NORMAL (2-way) or HIGH (3-way).",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "A user-friendly name for the VM cluster. This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain name associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "gi_version": {
        "description": "A valid software version of Oracle Grid Infrastructure (GI). To get the list of valid values, use the ListGiVersions operation and specify the shape of the Exadata infrastructure. Example: 19.0.0.0 This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gi_version_computed": {
        "computed": true,
        "description": "A complete software version of Oracle Grid Infrastructure (GI).",
        "description_kind": "plain",
        "type": "string"
      },
      "hostname_prefix": {
        "description": "The host name prefix for the VM cluster. Constraints: - Can't be \"localhost\" or \"hostname\". - Can't contain \"-version\". - The maximum length of the combined hostname and domain is 63 characters. - The hostname must be unique within the subnet. This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hostname_prefix_computed": {
        "computed": true,
        "description": "The host name for the VM cluster. Constraints: - Can't be \"localhost\" or \"hostname\". - Can't contain \"-version\". - The maximum length of the combined hostname and domain is 63 characters. - The hostname must be unique within the subnet. This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iorm_config_cache": {
        "computed": true,
        "description": "The Exadata IORM (I/O Resource Manager) configuration cache details for the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "db_plans": [
                "list",
                [
                  "object",
                  {
                    "db_name": "string",
                    "flash_cache_limit": "string",
                    "share": "number"
                  }
                ]
              ],
              "lifecycle_details": "string",
              "lifecycle_state": "string",
              "objective": "string"
            }
          ]
        ]
      },
      "is_local_backup_enabled": {
        "computed": true,
        "description": "Specifies whether to enable database backups to local Exadata storage for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_sparse_diskgroup_enabled": {
        "computed": true,
        "description": "Specifies whether to create a sparse disk group for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_update_history_entry_id": {
        "computed": true,
        "description": "The OCID of the most recent maintenance update history entry.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "The Oracle license model to apply to the VM cluster. Default: LICENSE_INCLUDED. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_port": {
        "computed": true,
        "description": "The listener port number configured on the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GBs), to allocate for the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_count": {
        "computed": true,
        "description": "The total number of nodes in the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the VM cluster resource in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID (Oracle Cloud Identifier) of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "description": "The unique identifier of the ODB network for the VM cluster. This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "The percentage of progress made on the current operation for the VM cluster.",
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
      "scan_dns_name": {
        "computed": true,
        "description": "The fully qualified domain name (FQDN) for the SCAN IP addresses associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_dns_record_id": {
        "computed": true,
        "description": "The OCID of the DNS record for the SCAN IPs linked to the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_ip_ids": {
        "computed": true,
        "description": "The list of OCIDs for SCAN IP addresses associated with the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scan_listener_port_tcp": {
        "computed": true,
        "description": "The port number for TCP connections to the single client access name (SCAN) listener. Valid values: 1024–8999 with the following exceptions: 2484 , 6100 , 6200 , 7060, 7070 , 7085 , and 7879Default: 1521. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shape": {
        "computed": true,
        "description": "The hardware model name of the Exadata infrastructure running the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssh_public_keys": {
        "description": "The public key portion of one or more key pairs used for SSH access to the VM cluster. This member is required. Changing this will create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The current lifecycle status of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information regarding the current status of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_size_in_gbs": {
        "computed": true,
        "description": "The local node storage allocated to the VM cluster, in gigabytes (GB).",
        "description_kind": "plain",
        "type": "number"
      },
      "system_version": {
        "computed": true,
        "description": "The operating system version of the image chosen for the VM cluster.",
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
      "timezone": {
        "computed": true,
        "description": "The configured time zone of the VM cluster. Changing this will create a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vip_ids": {
        "computed": true,
        "description": "The virtual IP (VIP) addresses assigned to the VM cluster. CRS assigns one VIP per node for failover support.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "data_collection_options": {
        "block": {
          "attributes": {
            "is_diagnostics_events_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "is_health_monitoring_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "is_incident_logs_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The set of preferences for the various diagnostic collection options for the VM cluster. Changing this will create a new resource.",
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

func AwsOdbCloudVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudVmCluster), &result)
	return &result
}
