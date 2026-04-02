package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerTrainingJob = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_model_packages_on_destroy": {
        "description": "Whether to delete model packages in the configured model package group when destroying the training job.",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
      },
      "delete_vpc_enis_on_destroy": {
        "description": "Whether to delete detached VPC ENIs that SageMaker may leave behind when destroying the training job.",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
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
      "hyper_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
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
      "training_job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "algorithm_specification": {
        "block": {
          "attributes": {
            "algorithm_name": {
              "description": "Name or ARN of a SageMaker algorithm resource. Exactly one of ` + "`" + `algorithm_name` + "`" + ` or ` + "`" + `training_image` + "`" + ` must be set.",
              "description_kind": "markdown",
              "optional": true,
              "type": "string"
            },
            "container_arguments": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "container_entrypoint": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enable_sagemaker_metrics_time_series": {
              "computed": true,
              "description": "Whether SageMaker AI should publish time-series metrics. SageMaker enables this automatically for built-in algorithms, supported prebuilt images, and jobs with explicit ` + "`" + `metric_definitions` + "`" + `.",
              "description_kind": "markdown",
              "optional": true,
              "type": "bool"
            },
            "training_image": {
              "description": "Registry path of the training image. Exactly one of ` + "`" + `algorithm_name` + "`" + ` or ` + "`" + `training_image` + "`" + ` must be set. Use ` + "`" + `metric_definitions` + "`" + ` only when you need to extract custom metrics from your own training container logs.",
              "description_kind": "markdown",
              "optional": true,
              "type": "string"
            },
            "training_input_mode": {
              "description_kind": "plain",
              "optional": true,
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
                "description": "Metric definitions used to extract custom metrics from training container logs. SageMaker may still return built-in metric definitions for built-in algorithms or supported prebuilt images even when this block is omitted.",
                "description_kind": "markdown"
              },
              "nesting_mode": "list"
            },
            "training_image_config": {
              "block": {
                "attributes": {
                  "training_repository_access_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "training_repository_auth_config": {
                    "block": {
                      "attributes": {
                        "training_repository_credentials_provider_arn": {
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
      "debug_hook_config": {
        "block": {
          "attributes": {
            "hook_parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "local_path": {
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
          "block_types": {
            "collection_configurations": {
              "block": {
                "attributes": {
                  "collection_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "collection_parameters": {
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
        "nesting_mode": "list"
      },
      "debug_rule_configurations": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_configuration_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_evaluator_image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "s3_output_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_size_in_gb": {
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
      "experiment_config": {
        "block": {
          "attributes": {
            "experiment_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "run_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trial_component_display_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trial_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "infra_check_config": {
        "block": {
          "attributes": {
            "enable_infra_check": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "input_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "record_wrapper_type": {
              "computed": true,
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
                            "list",
                            "string"
                          ]
                        },
                        "instance_group_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
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
      "mlflow_config": {
        "block": {
          "attributes": {
            "mlflow_experiment_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mlflow_resource_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mlflow_run_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "model_package_config": {
        "block": {
          "attributes": {
            "model_package_group_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_model_package_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
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
      "profiler_config": {
        "block": {
          "attributes": {
            "disable_profiler": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "profiling_interval_in_milliseconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "profiling_parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "s3_output_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "profiler_rule_configurations": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_configuration_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_evaluator_image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "s3_output_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_size_in_gb": {
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
      "remote_debug_config": {
        "block": {
          "attributes": {
            "enable_remote_debug": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "keep_alive_period_in_seconds": {
              "computed": true,
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
              "computed": true,
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
                    "optional": true,
                    "type": "number"
                  },
                  "instance_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
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
      "retry_strategy": {
        "block": {
          "attributes": {
            "maximum_retry_attempts": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "serverless_job_config": {
        "block": {
          "attributes": {
            "accept_eula": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "base_model_arn": {
              "description": "Base model ARN in SageMaker Public Hub. SageMaker always selects the latest version of the provided model.",
              "description_kind": "markdown",
              "required": true,
              "type": "string"
            },
            "customization_technique": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "evaluation_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "evaluator_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "job_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peft": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "session_chaining_config": {
        "block": {
          "attributes": {
            "enable_session_tag_chaining": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_runtime_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_wait_time_in_seconds": {
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
      "tensor_board_output_config": {
        "block": {
          "attributes": {
            "local_path": {
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
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
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

func AwsSagemakerTrainingJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerTrainingJob), &result)
	return &result
}
