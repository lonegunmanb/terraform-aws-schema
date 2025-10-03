package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbDbServer = `{
  "block": {
    "attributes": {
      "autonomous_virtual_machine_ids": {
        "computed": true,
        "description": "The list of unique identifiers for the Autonomous VMs associated with this database server.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "autonomous_vm_cluster_ids": {
        "computed": true,
        "description": "The OCID of the autonomous VM clusters that are associated with the database server.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cloud_exadata_infrastructure_id": {
        "description": "The identifier of the database server to retrieve information about.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": " The compute model of the database server.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "The number of CPU cores enabled on the database server.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the database server was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The allocated local node storage in GBs on the database server.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_server_patching_details": {
        "computed": true,
        "description": "The scheduling details for the quarterly maintenance window. Patching and\nsystem updates take place during the maintenance window.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "estimated_patch_duration": "number",
              "patching_status": "string",
              "time_patching_ended": "string",
              "time_patching_started": "string"
            }
          ]
        ]
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the database server.",
        "description_kind": "plain",
        "type": "string"
      },
      "exadata_infrastructure_id": {
        "computed": true,
        "description": "The exadata infrastructure ID of the database server.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "The identifier of the the database server.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores available.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_db_node_storage_in_gbs": {
        "computed": true,
        "description": "The total local node storage available in GBs.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_memory_in_gbs": {
        "computed": true,
        "description": "The total memory available in GBs.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The allocated memory in GBs on the database server.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the database server to retrieve information about.",
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
      "shape": {
        "computed": true,
        "description": "The shape of the database server. The shape determines the amount of CPU, storage, and memory resources available.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the database server.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the current status of the database server.",
        "description_kind": "plain",
        "type": "string"
      },
      "vm_cluster_ids": {
        "computed": true,
        "description": "The OCID of the VM clusters that are associated with the database server.",
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

func AwsOdbDbServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbDbServer), &result)
	return &result
}
