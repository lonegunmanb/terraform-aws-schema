package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDynamodbTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "billing_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hash_key": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "range_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_date_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_source_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_source_table_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_to_latest_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stream_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stream_label": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_view_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_class": {
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
      "write_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "attribute": {
        "block": {
          "attributes": {
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
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "global_secondary_index": {
        "block": {
          "attributes": {
            "hash_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "non_key_attributes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "projection_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "range_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read_capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "write_capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "on_demand_throughput": {
              "block": {
                "attributes": {
                  "max_read_request_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_write_request_units": {
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
        "nesting_mode": "set"
      },
      "import_table": {
        "block": {
          "attributes": {
            "input_compression_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "input_format": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "input_format_options": {
              "block": {
                "block_types": {
                  "csv": {
                    "block": {
                      "attributes": {
                        "delimiter": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "header_list": {
                          "description_kind": "plain",
                          "optional": true,
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
            "s3_bucket_source": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "bucket_owner": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_prefix": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "local_secondary_index": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "non_key_attributes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "projection_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "range_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "on_demand_throughput": {
        "block": {
          "attributes": {
            "max_read_request_units": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_write_request_units": {
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
      "point_in_time_recovery": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "recovery_period_in_days": {
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
      "replica": {
        "block": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "consistency_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "deletion_protection_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "point_in_time_recovery": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "propagate_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "region_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "stream_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stream_label": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "server_side_encryption": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "kms_key_arn": {
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
      "ttl": {
        "block": {
          "attributes": {
            "attribute_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsDynamodbTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDynamodbTable), &result)
	return &result
}
