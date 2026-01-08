package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsCapacityProvider = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
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
      }
    },
    "block_types": {
      "auto_scaling_group_provider": {
        "block": {
          "attributes": {
            "auto_scaling_group_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "managed_draining": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_termination_protection": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "managed_scaling": {
              "block": {
                "attributes": {
                  "instance_warmup_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minimum_scaling_step_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_capacity": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "managed_instances_provider": {
        "block": {
          "attributes": {
            "infrastructure_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "propagate_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "infrastructure_optimization": {
              "block": {
                "attributes": {
                  "scale_in_after": {
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
            "instance_launch_template": {
              "block": {
                "attributes": {
                  "capacity_option_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ec2_instance_profile_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "monitoring": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
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
                  "network_configuration": {
                    "block": {
                      "attributes": {
                        "security_groups": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "subnets": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "storage_configuration": {
                    "block": {
                      "attributes": {
                        "storage_size_gib": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
              "min_items": 1,
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
  "version": 0
}`

func AwsEcsCapacityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsCapacityProvider), &result)
	return &result
}
