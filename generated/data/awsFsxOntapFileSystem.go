package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxOntapFileSystem = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "automatic_backup_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "daily_automatic_backup_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_iops_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "iops": "number",
              "mode": "string"
            }
          ]
        ]
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_ip_address_range": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "intercluster": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "ha_pairs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "storage_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "throughput_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "throughput_capacity_per_ha_pair": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "weekly_maintenance_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsFsxOntapFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxOntapFileSystem), &result)
	return &result
}
