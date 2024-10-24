package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2InstanceType = `{
  "block": {
    "attributes": {
      "auto_recovery_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "bare_metal": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "burstable_performance_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "current_generation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "dedicated_hosts_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "default_cores": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "default_threads_per_core": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "default_vcpus": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_encryption_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_nvme_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_optimized_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_performance_baseline_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_performance_baseline_iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_performance_baseline_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_performance_maximum_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_performance_maximum_iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ebs_performance_maximum_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "efa_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ena_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_in_transit_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "fpgas": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "count": "number",
              "manufacturer": "string",
              "memory_size": "number",
              "name": "string"
            }
          ]
        ]
      },
      "free_tier_eligible": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "gpus": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "count": "number",
              "manufacturer": "string",
              "memory_size": "number",
              "name": "string"
            }
          ]
        ]
      },
      "hibernation_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hypervisor": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inference_accelerators": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "count": "number",
              "manufacturer": "string",
              "name": "string"
            }
          ]
        ]
      },
      "instance_disks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "count": "number",
              "size": "number",
              "type": "string"
            }
          ]
        ]
      },
      "instance_storage_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv6_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "maximum_ipv4_addresses_per_interface": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "maximum_ipv6_addresses_per_interface": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "maximum_network_cards": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "maximum_network_interfaces": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "network_performance": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "supported_architectures": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_placement_strategies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_root_device_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_usages_classes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_virtualization_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sustained_clock_speed": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_fpga_memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_gpu_memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_instance_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "valid_cores": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "number"
        ]
      },
      "valid_threads_per_core": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "number"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AwsEc2InstanceTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2InstanceType), &result)
	return &result
}
