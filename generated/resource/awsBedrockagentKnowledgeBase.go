package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentKnowledgeBase = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reasons": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "knowledge_base_configuration": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "vector_knowledge_base_configuration": {
              "block": {
                "attributes": {
                  "embedding_model_arn": {
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
      "storage_configuration": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "opensearch_serverless_configuration": {
              "block": {
                "attributes": {
                  "collection_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "vector_index_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "field_mapping": {
                    "block": {
                      "attributes": {
                        "metadata_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vector_field": {
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
            "pinecone_configuration": {
              "block": {
                "attributes": {
                  "connection_string": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "credentials_secret_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "field_mapping": {
                    "block": {
                      "attributes": {
                        "metadata_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
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
            "rds_configuration": {
              "block": {
                "attributes": {
                  "credentials_secret_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "field_mapping": {
                    "block": {
                      "attributes": {
                        "metadata_field": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "primary_key_field": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "text_field": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "vector_field": {
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
            "redis_enterprise_cloud_configuration": {
              "block": {
                "attributes": {
                  "credentials_secret_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "vector_index_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "field_mapping": {
                    "block": {
                      "attributes": {
                        "metadata_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vector_field": {
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

func AwsBedrockagentKnowledgeBaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentKnowledgeBase), &result)
	return &result
}
