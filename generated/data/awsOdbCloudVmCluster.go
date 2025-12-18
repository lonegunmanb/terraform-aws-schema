package data

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
      "cloud_exadata_infrastructure_arn": {
        "computed": true,
        "description": "The ARN of the Cloud Exadata Infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "computed": true,
        "description": "The ID of the Cloud Exadata Infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the Grid Infrastructure (GI) cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": "The OCI model compute model used when you create or clone an instance: ECPU or\nOCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on\nthe number of cores elastically allocated from a pool of compute and storage\nservers. An OCPU is a legacy physical measure of compute resources. OCPUs are\nbased on the physical core of a processor with hyper-threading enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "The number of CPU cores enabled on the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The time when the VM cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_collection_options": {
        "computed": true,
        "description": "The set of diagnostic collection options enabled for the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_diagnostics_events_enabled": "bool",
              "is_health_monitoring_enabled": "bool",
              "is_incident_logs_enabled": "bool"
            }
          ]
        ]
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description": "The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_servers": {
        "computed": true,
        "description": "The list of database servers for the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "disk_redundancy": {
        "computed": true,
        "description": "The type of redundancy configured for the VM cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain name of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "gi_version": {
        "computed": true,
        "description": "he software version of the Oracle Grid Infrastructure (GI) for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "hostname_prefix_computed": {
        "computed": true,
        "description": "The computed hostname prefix for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "The unique identifier of the VM cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "iorm_config_cache": {
        "computed": true,
        "description": "The ExadataIormConfig cache details for the VM cluster.",
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
        "description": "Indicates whether database backups to local Exadata storage is enabled for the VM cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_sparse_disk_group_enabled": {
        "computed": true,
        "description": "Indicates whether the VM cluster is configured with a sparse disk group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "last_update_history_entry_id": {
        "computed": true,
        "description": "The Oracle Cloud ID (OCID) of the last maintenance update history entry.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "The Oracle license model applied to the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "listener_port": {
        "computed": true,
        "description": "The port number configured for the listener on the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "node_count": {
        "computed": true,
        "description": "The number of nodes in the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI Resource Anchor.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the VM cluster in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "The ARN of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The ID of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "The amount of progress made on the current operation on the VM cluster,expressed as a percentage.",
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
        "description": "The FQDN of the DNS record for the Single Client Access Name (SCAN) IP\n addresses that are associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_dns_record_id": {
        "computed": true,
        "description": "The OCID of the DNS record for the SCAN IP addresses that are associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_ip_ids": {
        "computed": true,
        "description": "The OCID of the SCAN IP addresses that are associated with the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "shape": {
        "computed": true,
        "description": "The hardware model name of the Exadata infrastructure that's running the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssh_public_keys": {
        "computed": true,
        "description": "he public key portion of one or more key pairs used for SSH access to the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The status of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the status of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster.",
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
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "timezone": {
        "computed": true,
        "description": "The time zone of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "vip_ids": {
        "computed": true,
        "description": "The virtual IP (VIP) addresses that are associated with the VM cluster.\nOracle's Cluster Ready Services (CRS) creates and maintains one VIP address for\neach node in the VM cluster to enable failover. If one node fails, the VIP is\nreassigned to another active node in the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
