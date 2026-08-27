package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreMemoryStrategy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_execution_role_arn": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "memory_strategy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_templates": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "namespaces": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "consolidation": {
              "block": {
                "attributes": {
                  "append_to_prompt": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "extraction": {
              "block": {
                "attributes": {
                  "append_to_prompt": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "reflection": {
              "block": {
                "attributes": {
                  "append_to_prompt": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace_templates": {
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
            "self_managed_configuration": {
              "block": {
                "attributes": {
                  "historical_context_window_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "trigger_conditions_actual": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "message_based_trigger": [
                            "list",
                            [
                              "object",
                              {
                                "message_count": "number"
                              }
                            ]
                          ],
                          "time_based_trigger": [
                            "list",
                            [
                              "object",
                              {
                                "idle_session_timeout": "number"
                              }
                            ]
                          ],
                          "token_based_trigger": [
                            "list",
                            [
                              "object",
                              {
                                "token_count": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "invocation_configuration": {
                    "block": {
                      "attributes": {
                        "payload_delivery_bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "topic_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "trigger_conditions": {
                    "block": {
                      "block_types": {
                        "message_based_trigger": {
                          "block": {
                            "attributes": {
                              "message_count": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "time_based_trigger": {
                          "block": {
                            "attributes": {
                              "idle_session_timeout": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "token_based_trigger": {
                          "block": {
                            "attributes": {
                              "token_count": {
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
      "memory_record_schema": {
        "block": {
          "block_types": {
            "metadata_schema": {
              "block": {
                "attributes": {
                  "extraction_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "extraction_config": {
                    "block": {
                      "block_types": {
                        "llm_extraction_config": {
                          "block": {
                            "attributes": {
                              "definition": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "llm_extraction_instruction": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "validation": {
                                "block": {
                                  "block_types": {
                                    "number_validation": {
                                      "block": {
                                        "attributes": {
                                          "max_value": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "min_value": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "string_list_validation": {
                                      "block": {
                                        "attributes": {
                                          "allowed_values": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "max_items": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "string_validation": {
                                      "block": {
                                        "attributes": {
                                          "allowed_values": {
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
      "reflection_configuration": {
        "block": {
          "attributes": {
            "namespace_templates": {
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

func AwsBedrockagentcoreMemoryStrategySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreMemoryStrategy), &result)
	return &result
}
