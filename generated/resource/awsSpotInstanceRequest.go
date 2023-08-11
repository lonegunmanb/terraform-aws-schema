package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSpotInstanceRequest = `{
  "block": {
    "attributes": {
      "ami": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associate_public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_duration_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cpu_core_count": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cpu_threads_per_core": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disable_api_stop": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_api_termination": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "get_password_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hibernation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_resource_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_instance_profile": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_initiated_shutdown_behavior": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_interruption_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_address_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv6_addresses": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "key_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitoring": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "password_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "placement_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "placement_partition_number": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "primary_network_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_private_ips": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_dest_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "spot_bid_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "spot_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "spot_price": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_request_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "spot_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data_base64": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data_replace_on_change": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "valid_from": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "valid_until": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "wait_for_fulfillment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "capacity_reservation_specification": {
        "block": {
          "attributes": {
            "capacity_reservation_preference": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "capacity_reservation_target": {
              "block": {
                "attributes": {
                  "capacity_reservation_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "capacity_reservation_resource_group_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cpu_options": {
        "block": {
          "attributes": {
            "amd_sev_snp": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "core_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "threads_per_core": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "credit_specification": {
        "block": {
          "attributes": {
            "cpu_credits": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ebs_block_device": {
        "block": {
          "attributes": {
            "delete_on_termination": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "encrypted": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "snapshot_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
            "throughput": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volume_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "enclave_options": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ephemeral_block_device": {
        "block": {
          "attributes": {
            "device_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "no_device": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "virtual_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "launch_template": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_options": {
        "block": {
          "attributes": {
            "auto_recovery": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "metadata_options": {
        "block": {
          "attributes": {
            "http_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_protocol_ipv6": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_put_response_hop_limit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_tokens": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_metadata_tags": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "delete_on_termination": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_index": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "network_card_index": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "network_interface_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "private_dns_name_options": {
        "block": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "root_block_device": {
        "block": {
          "attributes": {
            "delete_on_termination": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "encrypted": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
            "throughput": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volume_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsSpotInstanceRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSpotInstanceRequest), &result)
	return &result
}
