package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerMonitoringSchedule = `{
  "block": {
    "attributes": {
      "arn": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "monitoring_schedule_config": {
        "block": {
          "attributes": {
            "monitoring_job_definition_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "monitoring_job_definition": {
              "block": {
                "attributes": {
                  "environment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "baseline": {
                    "block": {
                      "attributes": {
                        "baselining_job_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "constraints_resource": {
                          "block": {
                            "attributes": {
                              "s3_uri": {
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
                        "statistics_resource": {
                          "block": {
                            "attributes": {
                              "s3_uri": {
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
                  "monitoring_app_specification": {
                    "block": {
                      "attributes": {
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
                        "image_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "post_analytics_processor_source_uri": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "record_preprocessor_source_uri": {
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
                  "monitoring_inputs": {
                    "block": {
                      "block_types": {
                        "batch_transform_input": {
                          "block": {
                            "attributes": {
                              "data_captured_destination_s3_uri": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "end_time_offset": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exclude_features_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "features_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "inference_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "local_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "probability_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "probability_threshold_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "s3_data_distribution_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_input_mode": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_time_offset": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "dataset_format": {
                                "block": {
                                  "block_types": {
                                    "csv": {
                                      "block": {
                                        "attributes": {
                                          "header": {
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
                                    "json": {
                                      "block": {
                                        "attributes": {
                                          "line": {
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
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "endpoint_input": {
                          "block": {
                            "attributes": {
                              "end_time_offset": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "endpoint_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "exclude_features_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "features_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "inference_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "local_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "probability_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "probability_threshold_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "s3_data_distribution_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_input_mode": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_time_offset": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "monitoring_output_config": {
                    "block": {
                      "attributes": {
                        "kms_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "monitoring_outputs": {
                          "block": {
                            "block_types": {
                              "s3_output": {
                                "block": {
                                  "attributes": {
                                    "local_path": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "s3_upload_mode": {
                                      "computed": true,
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
                                "max_items": 1,
                                "min_items": 1,
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "monitoring_resources": {
                    "block": {
                      "block_types": {
                        "cluster_config": {
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
                              "volume_kms_key_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_size_in_gb": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "network_config": {
                    "block": {
                      "attributes": {
                        "enable_inter_container_traffic_encryption": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_network_isolation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
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
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "stopping_condition": {
                    "block": {
                      "attributes": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule_config": {
              "block": {
                "attributes": {
                  "schedule_expression": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerMonitoringScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerMonitoringSchedule), &result)
	return &result
}
