package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrSupportedInstanceTypes = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "release_label": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "supported_instance_types": {
        "block": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ebs_optimized_available": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "ebs_optimized_by_default": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "ebs_storage_only": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "instance_family_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "is_64_bits_only": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "memory_gb": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "number_of_disks": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "storage_gb": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vcpu": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEmrSupportedInstanceTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrSupportedInstanceTypes), &result)
	return &result
}
