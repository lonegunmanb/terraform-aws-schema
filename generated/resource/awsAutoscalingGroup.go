package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAutoscalingGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "capacity_rebalance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "context": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_cooldown": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_instance_warmup": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "desired_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "desired_capacity_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled_metrics": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_delete_warm_pool": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_check_grace_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_type": {
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
      "ignore_failed_scaling_activities": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "launch_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "max_instance_lifetime": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "metrics_granularity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_elb_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_size": {
        "description_kind": "plain",
        "required": true,
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
      "placement_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "predicted_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "protect_from_scale_in": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "service_linked_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suspended_processes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_group_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "termination_policies": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "wait_for_capacity_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait_for_elb_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "warm_pool_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "availability_zone_distribution": {
        "block": {
          "attributes": {
            "capacity_distribution_strategy": {
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
      "capacity_reservation_specification": {
        "block": {
          "attributes": {
            "capacity_reservation_preference": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "capacity_reservation_target": {
              "block": {
                "attributes": {
                  "capacity_reservation_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "capacity_reservation_resource_group_arns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "initial_lifecycle_hook": {
        "block": {
          "attributes": {
            "default_result": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "heartbeat_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "lifecycle_transition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "notification_metadata": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notification_target_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "instance_maintenance_policy": {
        "block": {
          "attributes": {
            "max_healthy_percentage": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_healthy_percentage": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "instance_refresh": {
        "block": {
          "attributes": {
            "strategy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "triggers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "preferences": {
              "block": {
                "attributes": {
                  "auto_rollback": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "checkpoint_delay": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "checkpoint_percentages": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "instance_warmup": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_healthy_percentage": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_healthy_percentage": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scale_in_protected_instances": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "skip_matching": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "standby_instances": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "alarm_specification": {
                    "block": {
                      "attributes": {
                        "alarms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
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
        "max_items": 1,
        "nesting_mode": "list"
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
      "mixed_instances_policy": {
        "block": {
          "block_types": {
            "instances_distribution": {
              "block": {
                "attributes": {
                  "on_demand_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_demand_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "on_demand_percentage_above_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_instance_pools": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_max_price": {
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
            "launch_template": {
              "block": {
                "block_types": {
                  "launch_template_specification": {
                    "block": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "override": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weighted_capacity": {
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
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
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
                        "launch_template_specification": {
                          "block": {
                            "attributes": {
                              "launch_template_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "launch_template_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "version": {
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
      "tag": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "propagate_at_launch": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "traffic_source": {
        "block": {
          "attributes": {
            "identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "warm_pool": {
        "block": {
          "attributes": {
            "max_group_prepared_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pool_state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "instance_reuse_policy": {
              "block": {
                "attributes": {
                  "reuse_on_scale_in": {
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
  "version": 1
}`

func AwsAutoscalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAutoscalingGroup), &result)
	return &result
}
