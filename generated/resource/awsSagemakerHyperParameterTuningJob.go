package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerHyperParameterTuningJob = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
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
      "status": {
        "computed": true,
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
      }
    },
    "block_types": {
      "autotune": {
        "block": {
          "attributes": {
            "mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "config": {
        "block": {
          "attributes": {
            "random_seed": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "strategy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "training_job_early_stopping_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "objective": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "parameter_ranges": {
              "block": {
                "block_types": {
                  "auto_parameters": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_hint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "categorical_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "continuous_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "integer_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "resource_limits": {
              "block": {
                "attributes": {
                  "max_number_of_training_jobs": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_parallel_training_jobs": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "max_runtime_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "strategy_config": {
              "block": {
                "block_types": {
                  "hyperband_strategy_config": {
                    "block": {
                      "attributes": {
                        "max_resource": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_resource": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "list"
            },
            "tuning_job_completion_criteria": {
              "block": {
                "attributes": {
                  "target_objective_metric_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "best_objective_not_improving": {
                    "block": {
                      "attributes": {
                        "max_number_of_training_jobs_not_improving": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "convergence_detected": {
                    "block": {
                      "attributes": {
                        "complete_on_convergence": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "training_job_definition": {
        "block": {
          "attributes": {
            "definition_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_inter_container_traffic_encryption": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_managed_spot_training": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_network_isolation": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "retry_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "maximum_retry_attempts": "number"
                  }
                ]
              ]
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "static_hyper_parameters": {
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
            "algorithm_specification": {
              "block": {
                "attributes": {
                  "algorithm_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "training_image": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "training_input_mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "metric_definitions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "regex": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "checkpoint_config": {
              "block": {
                "attributes": {
                  "local_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "hyper_parameter_ranges": {
              "block": {
                "block_types": {
                  "auto_parameters": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_hint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "categorical_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "continuous_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "integer_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "hyper_parameter_tuning_resource_config": {
              "block": {
                "attributes": {
                  "allocation_strategy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "instance_configs": {
                    "block": {
                      "attributes": {
                        "instance_count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume_size_in_gb": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "list"
            },
            "input_data_config": {
              "block": {
                "attributes": {
                  "channel_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "compression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "input_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_wrapper_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "data_source": {
                    "block": {
                      "block_types": {
                        "file_system_data_source": {
                          "block": {
                            "attributes": {
                              "directory_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_access_mode": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "s3_data_source": {
                          "block": {
                            "attributes": {
                              "attribute_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "instance_group_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "s3_data_distribution_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_data_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "s3_uri": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "hub_access_config": {
                                "block": {
                                  "attributes": {
                                    "hub_content_arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "model_access_config": {
                                "block": {
                                  "attributes": {
                                    "accept_eula": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
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
                    "nesting_mode": "list"
                  },
                  "shuffle_config": {
                    "block": {
                      "attributes": {
                        "seed": {
                          "description_kind": "plain",
                          "required": true,
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
              "nesting_mode": "list"
            },
            "output_data_config": {
              "block": {
                "attributes": {
                  "compression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_output_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "resource_config": {
              "block": {
                "attributes": {
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "keep_alive_period_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "training_plan_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "instance_groups": {
                    "block": {
                      "attributes": {
                        "instance_count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "instance_group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "instance_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "instance_placement_config": {
                    "block": {
                      "attributes": {
                        "enable_multiple_jobs": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "placement_specifications": {
                          "block": {
                            "attributes": {
                              "instance_count": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "ultra_server_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
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
              "nesting_mode": "list"
            },
            "stopping_condition": {
              "block": {
                "attributes": {
                  "max_pending_time_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_runtime_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_wait_time_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "tuning_objective": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_config": {
              "block": {
                "attributes": {
                  "security_group_ids": {
                    "description_kind": "plain",
                    "required": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "training_job_definitions": {
        "block": {
          "attributes": {
            "definition_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_inter_container_traffic_encryption": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_managed_spot_training": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_network_isolation": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "retry_strategy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "maximum_retry_attempts": "number"
                  }
                ]
              ]
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "static_hyper_parameters": {
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
            "algorithm_specification": {
              "block": {
                "attributes": {
                  "algorithm_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "training_image": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "training_input_mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "metric_definitions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "regex": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "checkpoint_config": {
              "block": {
                "attributes": {
                  "local_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "hyper_parameter_ranges": {
              "block": {
                "block_types": {
                  "auto_parameters": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_hint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "categorical_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "continuous_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "integer_parameter_ranges": {
                    "block": {
                      "attributes": {
                        "max_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "min_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scaling_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "hyper_parameter_tuning_resource_config": {
              "block": {
                "attributes": {
                  "allocation_strategy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "instance_configs": {
                    "block": {
                      "attributes": {
                        "instance_count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "volume_size_in_gb": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "list"
            },
            "input_data_config": {
              "block": {
                "attributes": {
                  "channel_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "compression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "input_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_wrapper_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "data_source": {
                    "block": {
                      "block_types": {
                        "file_system_data_source": {
                          "block": {
                            "attributes": {
                              "directory_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_access_mode": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "file_system_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "s3_data_source": {
                          "block": {
                            "attributes": {
                              "attribute_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "instance_group_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "s3_data_distribution_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_data_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "s3_uri": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "hub_access_config": {
                                "block": {
                                  "attributes": {
                                    "hub_content_arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "model_access_config": {
                                "block": {
                                  "attributes": {
                                    "accept_eula": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
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
                    "nesting_mode": "list"
                  },
                  "shuffle_config": {
                    "block": {
                      "attributes": {
                        "seed": {
                          "description_kind": "plain",
                          "required": true,
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
              "nesting_mode": "list"
            },
            "output_data_config": {
              "block": {
                "attributes": {
                  "compression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_output_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "resource_config": {
              "block": {
                "attributes": {
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "keep_alive_period_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "training_plan_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "instance_groups": {
                    "block": {
                      "attributes": {
                        "instance_count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "instance_group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "instance_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "instance_placement_config": {
                    "block": {
                      "attributes": {
                        "enable_multiple_jobs": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "placement_specifications": {
                          "block": {
                            "attributes": {
                              "instance_count": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "ultra_server_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
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
              "nesting_mode": "list"
            },
            "stopping_condition": {
              "block": {
                "attributes": {
                  "max_pending_time_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_runtime_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_wait_time_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "tuning_objective": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_config": {
              "block": {
                "attributes": {
                  "security_group_ids": {
                    "description_kind": "plain",
                    "required": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "warm_start_config": {
        "block": {
          "attributes": {
            "warm_start_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "parent_hyper_parameter_tuning_jobs": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
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
  "version": 0
}`

func AwsSagemakerHyperParameterTuningJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerHyperParameterTuningJob), &result)
	return &result
}
