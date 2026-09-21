package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskChannel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_type": {
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
      "encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "iceberg_destination": {
        "block": {
          "attributes": {
            "append_only": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "compression_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_freshness_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_execution_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "catalog": {
              "block": {
                "attributes": {
                  "catalog_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warehouse_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "dead_letter_queue_s3": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "destination_table": {
              "block": {
                "attributes": {
                  "destination_database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "destination_table_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "partition_spec": {
                    "block": {
                      "attributes": {
                        "partition_strategy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "source": {
                          "block": {
                            "attributes": {
                              "source_name": {
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
            "schema_evolution": {
              "block": {
                "attributes": {
                  "enable_schema_evolution": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "table_creation": {
              "block": {
                "attributes": {
                  "enable_table_creation": {
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
      },
      "logging_info": {
        "block": {
          "block_types": {
            "cloudwatch_logs": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "log_group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "firehose": {
              "block": {
                "attributes": {
                  "delivery_stream": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "s3": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "prefix": {
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
      "s3_destination": {
        "block": {
          "attributes": {
            "data_freshness_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_execution_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "dead_letter_queue_s3": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "storage": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "compression_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "expected_bucket_owner": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_key_template": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_class": {
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
      "topic_configuration": {
        "block": {
          "attributes": {
            "topic_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "record_converter": {
              "block": {
                "attributes": {
                  "value_converter": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "record_schema": {
              "block": {
                "attributes": {
                  "gsr_arn": {
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

func AwsMskChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskChannel), &result)
	return &result
}
