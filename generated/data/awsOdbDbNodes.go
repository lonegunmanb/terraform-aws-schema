package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbDbNodes = `{
  "block": {
    "attributes": {
      "cloud_vm_cluster_id": {
        "description": "Id of the cloud VM cluster. The unique identifier of the VM cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_nodes": {
        "computed": true,
        "description": "The list of DB nodes along with their properties.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "additional_details": "string",
              "arn": "string",
              "backup_ip_id": "string",
              "backup_vnic2_id": "string",
              "backup_vnic_id": "string",
              "cpu_core_count": "number",
              "created_at": "string",
              "db_node_storage_size": "number",
              "db_server_id": "string",
              "db_system_id": "string",
              "fault_domain": "string",
              "host_ip_id": "string",
              "hostname": "string",
              "id": "string",
              "maintenance_type": "string",
              "memory_size": "number",
              "oci_resource_anchor_name": "string",
              "ocid": "string",
              "software_storage_size": "number",
              "status": "string",
              "status_reason": "string",
              "time_maintenance_window_end": "string",
              "time_maintenance_window_start": "string",
              "total_cpu_core_count": "number",
              "vnic2_id": "string",
              "vnic_id": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbDbNodesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbDbNodes), &result)
	return &result
}
