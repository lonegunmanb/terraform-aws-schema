package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbDbServers = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure_id": {
        "description": "The cloud exadata infrastructure ID. Mandatory field.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_servers": {
        "computed": true,
        "description": "List of database servers associated with cloud_exadata_infrastructure_id.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "autonomous_virtual_machine_ids": [
                "list",
                "string"
              ],
              "autonomous_vm_cluster_ids": [
                "list",
                "string"
              ],
              "compute_model": "string",
              "cpu_core_count": "number",
              "created_at": "string",
              "db_node_storage_size_in_gbs": "number",
              "db_server_patching_details": [
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
              ],
              "display_name": "string",
              "exadata_infrastructure_id": "string",
              "id": "string",
              "max_cpu_count": "number",
              "max_db_node_storage_in_gbs": "number",
              "max_memory_in_gbs": "number",
              "memory_size_in_gbs": "number",
              "oci_resource_anchor_name": "string",
              "ocid": "string",
              "shape": "string",
              "status": "string",
              "status_reason": "string",
              "vm_cluster_ids": [
                "list",
                "string"
              ]
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

func AwsOdbDbServersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbDbServers), &result)
	return &result
}
