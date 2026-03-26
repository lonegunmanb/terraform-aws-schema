package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerAlgorithm = `{
  "block": {
    "attributes": {
      "algorithm_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "algorithm_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "algorithm_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certify_for_marketplace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description_kind": "plain",
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
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "inference_specification": {
        "block": {
          "attributes": {
            "supported_content_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_realtime_inference_instance_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_response_mime_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_transform_instance_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "containers": {
              "block": {
                "attributes": {
                  "container_hostname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "environment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "framework": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "framework_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_digest": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_checkpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "model_data_etag": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_data_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "nearest_model_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "product_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "additional_s3_data_source": {
                    "block": {
                      "attributes": {
                        "compression_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "etag": {
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
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "base_model": {
                    "block": {
                      "attributes": {
                        "hub_content_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "hub_content_version": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "recipe_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "model_data_source": {
                    "block": {
                      "block_types": {
                        "s3_data_source": {
                          "block": {
                            "attributes": {
                              "compression_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "etag": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "manifest_etag": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "manifest_s3_uri": {
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
                                      "optional": true,
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
                                      "optional": true,
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
                  "model_input": {
                    "block": {
                      "attributes": {
                        "data_input_config": {
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
      "training_specification": {
        "block": {
          "attributes": {
            "supported_training_instance_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supports_distributed_training": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "training_image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "training_image_digest": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "additional_s3_data_source": {
              "block": {
                "attributes": {
                  "compression_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "etag": {
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
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
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
            },
            "supported_hyper_parameters": {
              "block": {
                "attributes": {
                  "default_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "is_tunable": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
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
                "block_types": {
                  "range": {
                    "block": {
                      "block_types": {
                        "categorical_parameter_range_specification": {
                          "block": {
                            "attributes": {
                              "values": {
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
                        },
                        "continuous_parameter_range_specification": {
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
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "integer_parameter_range_specification": {
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
            "supported_tuning_job_objective_metrics": {
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
            "training_channels": {
              "block": {
                "attributes": {
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "supported_compression_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "supported_content_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "supported_input_modes": {
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
        "nesting_mode": "list"
      },
      "validation_specification": {
        "block": {
          "attributes": {
            "validation_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "validation_profiles": {
              "block": {
                "attributes": {
                  "profile_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "training_job_definition": {
                    "block": {
                      "attributes": {
                        "hyper_parameters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "training_input_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
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
                                "computed": true,
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
                                                  "optional": true,
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
                                                  "optional": true,
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
                                "computed": true,
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
                                "computed": true,
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
                                "computed": true,
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
                  "transform_job_definition": {
                    "block": {
                      "attributes": {
                        "batch_strategy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "environment": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "max_concurrent_transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_payload_in_mb": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "transform_input": {
                          "block": {
                            "attributes": {
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
                              "split_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "data_source": {
                                "block": {
                                  "block_types": {
                                    "s3_data_source": {
                                      "block": {
                                        "attributes": {
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
                        "transform_output": {
                          "block": {
                            "attributes": {
                              "accept": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "assemble_with": {
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
                        "transform_resources": {
                          "block": {
                            "attributes": {
                              "instance_count": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "instance_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "transform_ami_version": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_kms_key_id": {
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

func AwsSagemakerAlgorithmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerAlgorithm), &result)
	return &result
}
