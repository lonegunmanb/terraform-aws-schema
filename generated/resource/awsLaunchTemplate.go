package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLaunchTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_api_stop": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_api_termination": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ebs_optimized": {
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
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_initiated_shutdown_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kernel_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_disk_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      "update_default_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "block_device_mappings": {
        "block": {
          "attributes": {
            "device_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_device": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "virtual_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ebs": {
              "block": {
                "attributes": {
                  "delete_on_termination": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encrypted": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "snapshot_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "throughput": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "volume_initialization_rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "core_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "threads_per_core": {
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
      "elastic_gpu_specifications": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "elastic_inference_accelerator": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "enclave_options": {
        "block": {
          "attributes": {
            "enabled": {
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
      "hibernation_options": {
        "block": {
          "attributes": {
            "configured": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "iam_instance_profile": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
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
      "instance_market_options": {
        "block": {
          "attributes": {
            "market_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "spot_options": {
              "block": {
                "attributes": {
                  "block_duration_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_interruption_behavior": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "valid_until": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "instance_requirements": {
        "block": {
          "attributes": {
            "accelerator_manufacturers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "accelerator_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "accelerator_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "allowed_instance_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "bare_metal": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "burstable_performance": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cpu_manufacturers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "excluded_instance_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "instance_generations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "local_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_storage_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "max_spot_price_as_percentage_of_optimal_on_demand_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "on_demand_max_price_percentage_over_lowest_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "require_hibernate_support": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "spot_max_price_percentage_over_lowest_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "accelerator_count": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "accelerator_total_memory_mib": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "baseline_ebs_bandwidth_mbps": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "memory_gib_per_vcpu": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "memory_mib": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "network_bandwidth_gbps": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "network_interface_count": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "total_local_storage_gb": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
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
            "vcpu_count": {
              "block": {
                "attributes": {
                  "max": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "license_specification": {
        "block": {
          "attributes": {
            "license_configuration_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "maintenance_options": {
        "block": {
          "attributes": {
            "auto_recovery": {
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_protocol_ipv6": {
              "computed": true,
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
      "monitoring": {
        "block": {
          "attributes": {
            "enabled": {
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
      "network_interfaces": {
        "block": {
          "attributes": {
            "associate_carrier_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "associate_public_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete_on_termination": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_index": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interface_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv4_address_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv4_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "ipv4_prefix_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv4_prefixes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "ipv6_address_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv6_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "ipv6_prefix_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv6_prefixes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "network_card_index": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "network_interface_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary_ipv6": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "connection_tracking_specification": {
              "block": {
                "attributes": {
                  "tcp_established_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "udp_stream_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "udp_timeout": {
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
            "ena_srd_specification": {
              "block": {
                "attributes": {
                  "ena_srd_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "ena_srd_udp_specification": {
                    "block": {
                      "attributes": {
                        "ena_srd_udp_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "placement": {
        "block": {
          "attributes": {
            "affinity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "availability_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_resource_group_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partition_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spread_domain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tenancy": {
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
      "private_dns_name_options": {
        "block": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_type": {
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
      "tag_specifications": {
        "block": {
          "attributes": {
            "resource_type": {
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

func AwsLaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLaunchTemplate), &result)
	return &result
}
