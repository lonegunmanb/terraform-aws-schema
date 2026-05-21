package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreHarness = `{
  "block": {
    "attributes": {
      "allowed_tools": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "agentcore_runtime_environment": [
                "list",
                [
                  "object",
                  {
                    "agent_runtime_arn": "string",
                    "agent_runtime_id": "string",
                    "agent_runtime_name": "string",
                    "filesystem_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "efs_access_point": [
                            "list",
                            [
                              "object",
                              {
                                "access_point_arn": "string",
                                "mount_path": "string"
                              }
                            ]
                          ],
                          "s3_files_access_point": [
                            "list",
                            [
                              "object",
                              {
                                "access_point_arn": "string",
                                "mount_path": "string"
                              }
                            ]
                          ],
                          "session_storage": [
                            "list",
                            [
                              "object",
                              {
                                "mount_path": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "lifecycle_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "idle_runtime_session_timeout": "number",
                          "max_lifetime": "number"
                        }
                      ]
                    ],
                    "network_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "network_mode": "string",
                          "network_mode_config": [
                            "list",
                            [
                              "object",
                              {
                                "security_groups": [
                                  "set",
                                  "string"
                                ],
                                "subnets": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "environment_variables": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "harness_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "harness_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_iterations": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_tokens": {
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
      "timeout_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "truncation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "config": [
                "list",
                [
                  "object",
                  {
                    "sliding_window": [
                      "list",
                      [
                        "object",
                        {
                          "messages_count": "number"
                        }
                      ]
                    ],
                    "summarization": [
                      "list",
                      [
                        "object",
                        {
                          "preserve_recent_messages": "number",
                          "summarization_system_prompt": "string",
                          "summary_ratio": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "strategy": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "authorizer_configuration": {
        "block": {
          "block_types": {
            "custom_jwt_authorizer": {
              "block": {
                "attributes": {
                  "allowed_audience": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "allowed_scopes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "discovery_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_claim": {
                    "block": {
                      "attributes": {
                        "inbound_token_claim_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "inbound_token_claim_value_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "authorizing_claim_match_value": {
                          "block": {
                            "attributes": {
                              "claim_match_operator": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "claim_match_value": {
                                "block": {
                                  "attributes": {
                                    "match_value_string": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "match_value_string_list": {
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
                    "nesting_mode": "set"
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
      "environment_artifact": {
        "block": {
          "block_types": {
            "container_configuration": {
              "block": {
                "attributes": {
                  "container_uri": {
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
      "memory": {
        "block": {
          "block_types": {
            "agentcore_memory_configuration": {
              "block": {
                "attributes": {
                  "actor_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "messages_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "retrieval_config": {
                    "block": {
                      "attributes": {
                        "map_block_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "relevance_score": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "strategy_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "top_k": {
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
      "model": {
        "block": {
          "block_types": {
            "bedrock_model_config": {
              "block": {
                "attributes": {
                  "max_tokens": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
            },
            "gemini_model_config": {
              "block": {
                "attributes": {
                  "api_key_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "max_tokens": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "temperature": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "top_k": {
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
            },
            "openai_model_config": {
              "block": {
                "attributes": {
                  "api_key_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "max_tokens": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "model_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
      "skill": {
        "block": {
          "attributes": {
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "system_prompt": {
        "block": {
          "attributes": {
            "text": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
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
      "tool": {
        "block": {
          "attributes": {
            "name": {
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
            "config": {
              "block": {
                "block_types": {
                  "agentcore_browser": {
                    "block": {
                      "attributes": {
                        "browser_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "agentcore_code_interpreter": {
                    "block": {
                      "attributes": {
                        "code_interpreter_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "agentcore_gateway": {
                    "block": {
                      "attributes": {
                        "gateway_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "outbound_auth": {
                          "block": {
                            "attributes": {
                              "aws_iam": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "none": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "oauth": {
                                "block": {
                                  "attributes": {
                                    "custom_parameters": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "default_return_url": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "grant_type": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "provider_arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "scopes": {
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
                  },
                  "inline_function": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "input_schema": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "remote_mcp": {
                    "block": {
                      "attributes": {
                        "headers": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "url": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
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
  "version": 0
}`

func AwsBedrockagentcoreHarnessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreHarness), &result)
	return &result
}
