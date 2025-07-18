package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentFlow = `{
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
      "customer_encryption_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "definition": {
        "block": {
          "block_types": {
            "connection": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "target": {
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
                  "configuration": {
                    "block": {
                      "block_types": {
                        "conditional": {
                          "block": {
                            "attributes": {
                              "condition": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "data": {
                          "block": {
                            "attributes": {
                              "source_output": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "target_input": {
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
            "node": {
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
                "block_types": {
                  "configuration": {
                    "block": {
                      "block_types": {
                        "agent": {
                          "block": {
                            "attributes": {
                              "agent_alias_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "collector": {
                          "block": {
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "condition": {
                          "block": {
                            "block_types": {
                              "condition": {
                                "block": {
                                  "attributes": {
                                    "expression": {
                                      "description_kind": "plain",
                                      "optional": true,
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
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "inline_code": {
                          "block": {
                            "attributes": {
                              "code": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "language": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "input": {
                          "block": {
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "iterator": {
                          "block": {
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "knowledge_base": {
                          "block": {
                            "attributes": {
                              "knowledge_base_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "model_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "number_of_results": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "guardrail_configuration": {
                                "block": {
                                  "attributes": {
                                    "guardrail_identifier": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "guardrail_version": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "inference_configuration": {
                                "block": {
                                  "block_types": {
                                    "text": {
                                      "block": {
                                        "attributes": {
                                          "max_tokens": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "stop_sequences": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "temperature": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "top_p": {
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
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "lambda_function": {
                          "block": {
                            "attributes": {
                              "lambda_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "lex": {
                          "block": {
                            "attributes": {
                              "bot_alias_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "locale_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "output": {
                          "block": {
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "prompt": {
                          "block": {
                            "block_types": {
                              "guardrail_configuration": {
                                "block": {
                                  "attributes": {
                                    "guardrail_identifier": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "guardrail_version": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "source_configuration": {
                                "block": {
                                  "block_types": {
                                    "inline": {
                                      "block": {
                                        "attributes": {
                                          "additional_model_request_fields": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "model_id": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "template_type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "inference_configuration": {
                                            "block": {
                                              "block_types": {
                                                "text": {
                                                  "block": {
                                                    "attributes": {
                                                      "max_tokens": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "stop_sequences": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": [
                                                          "list",
                                                          "string"
                                                        ]
                                                      },
                                                      "temperature": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "top_p": {
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
                                          "template_configuration": {
                                            "block": {
                                              "block_types": {
                                                "chat": {
                                                  "block": {
                                                    "block_types": {
                                                      "input_variable": {
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
                                                      },
                                                      "message": {
                                                        "block": {
                                                          "attributes": {
                                                            "role": {
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "block_types": {
                                                            "content": {
                                                              "block": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "description_kind": "plain",
                                                                    "optional": true,
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "block_types": {
                                                                  "cache_point": {
                                                                    "block": {
                                                                      "attributes": {
                                                                        "type": {
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
                                                      "system": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            }
                                                          },
                                                          "block_types": {
                                                            "cache_point": {
                                                              "block": {
                                                                "attributes": {
                                                                  "type": {
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
                                                      "tool_configuration": {
                                                        "block": {
                                                          "block_types": {
                                                            "tool": {
                                                              "block": {
                                                                "block_types": {
                                                                  "cache_point": {
                                                                    "block": {
                                                                      "attributes": {
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
                                                                  "tool_spec": {
                                                                    "block": {
                                                                      "attributes": {
                                                                        "description": {
                                                                          "description_kind": "plain",
                                                                          "optional": true,
                                                                          "type": "string"
                                                                        },
                                                                        "name": {
                                                                          "description_kind": "plain",
                                                                          "required": true,
                                                                          "type": "string"
                                                                        }
                                                                      },
                                                                      "block_types": {
                                                                        "input_schema": {
                                                                          "block": {
                                                                            "attributes": {
                                                                              "json": {
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
                                                            "tool_choice": {
                                                              "block": {
                                                                "block_types": {
                                                                  "any": {
                                                                    "block": {
                                                                      "description_kind": "plain"
                                                                    },
                                                                    "nesting_mode": "list"
                                                                  },
                                                                  "auto": {
                                                                    "block": {
                                                                      "description_kind": "plain"
                                                                    },
                                                                    "nesting_mode": "list"
                                                                  },
                                                                  "tool": {
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
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "text": {
                                                  "block": {
                                                    "attributes": {
                                                      "text": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "cache_point": {
                                                        "block": {
                                                          "attributes": {
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
                                                      "input_variable": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "resource": {
                                      "block": {
                                        "attributes": {
                                          "resource_arn": {
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
                        "retrieval": {
                          "block": {
                            "block_types": {
                              "service_configuration": {
                                "block": {
                                  "block_types": {
                                    "s3": {
                                      "block": {
                                        "attributes": {
                                          "bucket_name": {
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
                        "storage": {
                          "block": {
                            "block_types": {
                              "service_configuration": {
                                "block": {
                                  "block_types": {
                                    "s3": {
                                      "block": {
                                        "attributes": {
                                          "bucket_name": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "input": {
                    "block": {
                      "attributes": {
                        "category": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "output": {
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

func AwsBedrockagentFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentFlow), &result)
	return &result
}
