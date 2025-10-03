package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbDbNode = `{
  "block": {
    "attributes": {
      "additional_details": {
        "computed": true,
        "description": "Additional information about the planned maintenance.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_ip_id": {
        "computed": true,
        "description": "The Oracle Cloud ID (OCID) of the backup IP address that's associated with the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vnic2_id": {
        "computed": true,
        "description": "The OCID of the second backup VNIC.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vnic_id": {
        "computed": true,
        "description": "The OCID of the backup VNIC.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_vm_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "Number of CPU cores enabled on the DB node.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the DB node was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_server_id": {
        "computed": true,
        "description": "The unique identifier of the DB server that is associated with the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GBs), allocated on the DB node.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_system_id": {
        "computed": true,
        "description": "The OCID of the DB system.",
        "description_kind": "plain",
        "type": "string"
      },
      "fault_domain": {
        "computed": true,
        "description": "The name of the fault domain the instance is contained in.",
        "description_kind": "plain",
        "type": "string"
      },
      "floating_ip_address": {
        "computed": true,
        "description": "The floating IP address assigned to the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_ip_id": {
        "computed": true,
        "description": "The OCID of the host IP address that's associated with the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "hostname": {
        "computed": true,
        "description": "The host name for the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_type": {
        "computed": true,
        "description": "The type of database node maintenance. Either VMDB_REBOOT_MIGRATION or EXADBXS_REBOOT_MIGRATION.",
        "description_kind": "plain",
        "type": "string"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The allocated memory in GBs on the DB node.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "The private IP address assigned to the DB node.",
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
      "software_storage_size_in_gbs": {
        "computed": true,
        "description": "The size (in GB) of the block storage volume allocation for the DB system.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The current status of the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the status of the DB node.",
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_window_end": {
        "computed": true,
        "description": "End date and time of maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_window_start": {
        "computed": true,
        "description": "Start date and time of maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "total_cpu_core_count": {
        "computed": true,
        "description": "The total number of CPU cores reserved on the DB node.",
        "description_kind": "plain",
        "type": "number"
      },
      "vnic2_id": {
        "computed": true,
        "description": "The OCID of the second VNIC.",
        "description_kind": "plain",
        "type": "string"
      },
      "vnic_id": {
        "computed": true,
        "description": "The OCID of the VNIC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbDbNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbDbNode), &result)
	return &result
}
