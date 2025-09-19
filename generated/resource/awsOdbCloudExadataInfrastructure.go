package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudExadataInfrastructure = `{
  "block": {
    "attributes": {
      "activated_storage_count": {
        "computed": true,
        "description": "The number of storage servers requested for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "additional_storage_count": {
        "computed": true,
        "description": " The number of storage servers requested for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The name of the Availability Zone (AZ) where the Exadata infrastructure is located. Changing this will force terraform to create new resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "description": " The AZ ID of the AZ where the Exadata infrastructure is located. Changing this will force terraform to create new resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "available_storage_size_in_gbs": {
        "computed": true,
        "description": "The amount of available storage, in gigabytes (GB), for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "compute_count": {
        "computed": true,
        "description": " The number of compute instances that the Exadata infrastructure is located",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description": "The OCI model compute model used when you create or clone an\n  instance: ECPU or OCPU. An ECPU is an abstracted measure of\n compute resources. ECPUs are based on the number of cores\n elastically allocated from a pool of compute and storage servers.\n  An OCPU is a legacy physical measure of compute resources. OCPUs\n are based on the physical core of a processor with\n  hyper-threading enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores that are allocated to the Exadata infrastructure",
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
        "description": "The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure. Changing this will force terraform to create new resource",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The size of the Exadata infrastructure's data disk group, in terabytes (TB)",
        "description_kind": "plain",
        "type": "number"
      },
      "database_server_type": {
        "description": "The database server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The size of the Exadata infrastructure's local node storage, in gigabytes (GB)",
        "description_kind": "plain",
        "type": "number"
      },
      "db_server_version": {
        "computed": true,
        "description": "The software version of the database servers (dom0) in the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The user-friendly name for the Exadata infrastructure. Changing this will force terraform to create a new resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_maintenance_run_id": {
        "computed": true,
        "description": "The Oracle Cloud Identifier (OCID) of the last maintenance run for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "max_cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores available on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "max_data_storage_in_tbs": {
        "computed": true,
        "description": "The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "max_db_node_storage_size_in_gbs": {
        "computed": true,
        "description": "The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "max_memory_in_gbs": {
        "computed": true,
        "description": "The total amount of memory in gigabytes (GB) available on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "number"
      },
      "monthly_db_server_version": {
        "computed": true,
        "description": "The monthly software version of the database servers in the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "monthly_storage_server_version": {
        "computed": true,
        "description": "The monthly software version of the storage servers installed on the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "next_maintenance_run_id": {
        "computed": true,
        "description": "The OCID of the next maintenance run for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the Exadata infrastructure in OCI",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "The amount of progress made on the current operation on the Exadata infrastructure, expressed as a percentage",
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
        "description": "The model name of the Exadata infrastructure. Changing this will force terraform to create new resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the status of the Exadata infrastructure",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_count": {
        "computed": true,
        "description": "TThe number of storage servers that are activated for the Exadata infrastructure",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_server_type": {
        "description": "The storage server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_server_version": {
        "computed": true,
        "description": "The software version of the storage servers on the Exadata infrastructure.",
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
      "total_storage_size_in_gbs": {
        "computed": true,
        "description": "The total amount of storage, in gigabytes (GB), on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "maintenance_window": {
        "block": {
          "attributes": {
            "custom_action_timeout_in_mins": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "days_of_week": {
              "computed": true,
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "is_custom_action_timeout_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "lead_time_in_weeks": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "months": {
              "computed": true,
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
            "patching_mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "preference": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weeks_of_month": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            }
          },
          "description": " The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window ",
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

func AwsOdbCloudExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudExadataInfrastructure), &result)
	return &result
}
