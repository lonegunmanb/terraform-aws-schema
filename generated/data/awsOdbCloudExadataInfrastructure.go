package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudExadataInfrastructure = `{
  "block": {
    "attributes": {
      "activated_storage_count": {
        "computed": true,
        "description": "The number of storage servers requested for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "additional_storage_count": {
        "computed": true,
        "description": "The number of storage servers requested for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "he name of the Availability Zone (AZ) where the Exadata infrastructure is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the AZ where the Exadata infrastructure is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "available_storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of available storage, in gigabytes (GB), for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "compute_count": {
        "computed": true,
        "description": "The number of database servers for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description": "The OCI model compute model used when you create or clone an instance: ECPU or\nOCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on\nthe number of cores elastically allocated from a pool of compute and storage\nservers. An OCPU is a legacy physical measure of compute resources. OCPUs are\nbased on the physical core of a processor with hyper-threading enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores that are allocated to the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description": "The time when the Exadata infrastructure was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_contacts_to_send_to_oci": {
        "computed": true,
        "description": "The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "email": "string"
            }
          ]
        ]
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description": "The size of the Exadata infrastructure's data disk group, in terabytes (TB).",
        "description_kind": "plain",
        "type": "number"
      },
      "database_server_type": {
        "computed": true,
        "description": "The database server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The database server model type of the Exadata infrastructure. For the list of\nvalid model names, use the ListDbSystemShapes operation.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_server_version": {
        "computed": true,
        "description": "The version of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "The unique identifier of the Exadata infrastructure.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_maintenance_run_id": {
        "computed": true,
        "description": "The Oracle Cloud Identifier (OCID) of the last maintenance run for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": " The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window ",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_action_timeout_in_mins": "number",
              "days_of_week": [
                "set",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "hours_of_day": [
                "set",
                "number"
              ],
              "is_custom_action_timeout_enabled": "bool",
              "lead_time_in_weeks": "number",
              "months": [
                "set",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "patching_mode": "string",
              "preference": "string",
              "weeks_of_month": [
                "set",
                "number"
              ]
            }
          ]
        ]
      },
      "max_cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_data_storage_in_tbs": {
        "computed": true,
        "description": "The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_memory_in_gbs": {
        "computed": true,
        "description": "The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "monthly_db_server_version": {
        "computed": true,
        "description": "The monthly software version of the database servers installed on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "monthly_storage_server_version": {
        "computed": true,
        "description": "The monthly software version of the storage servers installed on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "next_maintenance_run_id": {
        "computed": true,
        "description": "The OCID of the next maintenance run for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the Exadata infrastructure in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the Exadata infrastructure in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "The amount of progress made on the current operation on the Exadata infrastructure expressed as a percentage.",
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
      "shape": {
        "computed": true,
        "description": "The model name of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the status of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_count": {
        "computed": true,
        "description": "he number of storage servers that are activated for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_server_type": {
        "computed": true,
        "description": "The storage server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_server_version": {
        "computed": true,
        "description": "The software version of the storage servers on the Exadata infrastructure.",
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
      "total_storage_size_in_gbs": {
        "computed": true,
        "description": "The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbCloudExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudExadataInfrastructure), &result)
	return &result
}
