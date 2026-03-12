package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueCatalogTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "owner": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
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
      "retention": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "table_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "view_expanded_text": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "view_original_text": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "open_table_format_input": {
        "block": {
          "block_types": {
            "iceberg_input": {
              "block": {
                "attributes": {
                  "metadata_operation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "iceberg_table_input": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "properties": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "partition_spec": {
                          "block": {
                            "attributes": {
                              "spec_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "fields": {
                                "block": {
                                  "attributes": {
                                    "field_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "source_id": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "transform": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "schema": {
                          "block": {
                            "attributes": {
                              "identifier_field_ids": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              },
                              "schema_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "fields": {
                                "block": {
                                  "attributes": {
                                    "doc": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "id": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "initial_default": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "required": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    },
                                    "type": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "write_default": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
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
                        "sort_order": {
                          "block": {
                            "attributes": {
                              "order_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "fields": {
                                "block": {
                                  "attributes": {
                                    "direction": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "null_order": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "source_id": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "transform": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
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
      "partition_index": {
        "block": {
          "attributes": {
            "index_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "index_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "keys": {
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
        "max_items": 3,
        "nesting_mode": "list"
      },
      "partition_keys": {
        "block": {
          "attributes": {
            "comment": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "storage_descriptor": {
        "block": {
          "attributes": {
            "additional_locations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "bucket_columns": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "compressed": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "input_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "number_of_buckets": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "output_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "stored_as_sub_directories": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "columns": {
              "block": {
                "attributes": {
                  "comment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "schema_reference": {
              "block": {
                "attributes": {
                  "schema_version_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schema_version_number": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "schema_id": {
                    "block": {
                      "attributes": {
                        "registry_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "schema_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "schema_name": {
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
            "ser_de_info": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "serialization_library": {
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
            "skewed_info": {
              "block": {
                "attributes": {
                  "skewed_column_names": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "skewed_column_value_location_maps": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "skewed_column_values": {
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
            },
            "sort_columns": {
              "block": {
                "attributes": {
                  "column": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sort_order": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target_table": {
        "block": {
          "attributes": {
            "catalog_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "database_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
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
      "view_definition": {
        "block": {
          "attributes": {
            "definer": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_protected": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "last_refresh_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refresh_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sub_object_version_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "sub_objects": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "view_version_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "view_version_token": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "representations": {
              "block": {
                "attributes": {
                  "dialect": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dialect_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "validation_connection": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "view_expanded_text": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "view_original_text": {
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
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueCatalogTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueCatalogTable), &result)
	return &result
}
