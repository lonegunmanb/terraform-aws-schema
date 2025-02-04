package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTimestreamqueryScheduledQuery = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_invocation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "previous_invocation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
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
      "error_report_configuration": {
        "block": {
          "block_types": {
            "s3_configuration": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "encryption_option": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "object_key_prefix": {
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
      "last_run_summary": {
        "block": {
          "attributes": {
            "failure_reason": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "invocation_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "run_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "trigger_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "error_report_location": {
              "block": {
                "block_types": {
                  "s3_report_location": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "object_key": {
                          "computed": true,
                          "description_kind": "plain",
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
            "execution_stats": {
              "block": {
                "attributes": {
                  "bytes_metered": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "cumulative_bytes_scanned": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "data_writes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "execution_time_in_millis": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "query_result_rows": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "records_ingested": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "query_insights_response": {
              "block": {
                "attributes": {
                  "output_bytes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "output_rows": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "query_table_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "block_types": {
                  "query_spatial_coverage": {
                    "block": {
                      "block_types": {
                        "max": {
                          "block": {
                            "attributes": {
                              "partition_key": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "table_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
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
                    "nesting_mode": "list"
                  },
                  "query_temporal_range": {
                    "block": {
                      "block_types": {
                        "max": {
                          "block": {
                            "attributes": {
                              "table_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
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
      "notification_configuration": {
        "block": {
          "block_types": {
            "sns_configuration": {
              "block": {
                "attributes": {
                  "topic_arn": {
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
      "recently_failed_runs": {
        "block": {
          "attributes": {
            "failure_reason": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "invocation_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "run_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "trigger_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "error_report_location": {
              "block": {
                "block_types": {
                  "s3_report_location": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "object_key": {
                          "computed": true,
                          "description_kind": "plain",
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
            "execution_stats": {
              "block": {
                "attributes": {
                  "bytes_metered": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "cumulative_bytes_scanned": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "data_writes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "execution_time_in_millis": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "query_result_rows": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "records_ingested": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "query_insights_response": {
              "block": {
                "attributes": {
                  "output_bytes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "output_rows": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "query_table_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "block_types": {
                  "query_spatial_coverage": {
                    "block": {
                      "block_types": {
                        "max": {
                          "block": {
                            "attributes": {
                              "partition_key": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "table_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
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
                    "nesting_mode": "list"
                  },
                  "query_temporal_range": {
                    "block": {
                      "block_types": {
                        "max": {
                          "block": {
                            "attributes": {
                              "table_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
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
      "schedule_configuration": {
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
        "nesting_mode": "list"
      },
      "target_configuration": {
        "block": {
          "block_types": {
            "timestream_configuration": {
              "block": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "measure_name_column": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time_column": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "dimension_mapping": {
                    "block": {
                      "attributes": {
                        "dimension_value_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "mixed_measure_mapping": {
                    "block": {
                      "attributes": {
                        "measure_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "measure_value_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "source_column": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "target_measure_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "multi_measure_attribute_mapping": {
                          "block": {
                            "attributes": {
                              "measure_value_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source_column": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "target_multi_measure_attribute_name": {
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
                  "multi_measure_mappings": {
                    "block": {
                      "attributes": {
                        "target_multi_measure_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "multi_measure_attribute_mapping": {
                          "block": {
                            "attributes": {
                              "measure_value_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source_column": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "target_multi_measure_attribute_name": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTimestreamqueryScheduledQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTimestreamqueryScheduledQuery), &result)
	return &result
}
