package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreGatewayTarget = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_identifier": {
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
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "credential_provider_configuration": {
        "block": {
          "block_types": {
            "api_key": {
              "block": {
                "attributes": {
                  "credential_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "credential_parameter_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "credential_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "provider_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "gateway_iam_role": {
              "block": {
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
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
                  "provider_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "scopes": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "target_configuration": {
        "block": {
          "block_types": {
            "mcp": {
              "block": {
                "block_types": {
                  "lambda": {
                    "block": {
                      "attributes": {
                        "lambda_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "tool_schema": {
                          "block": {
                            "block_types": {
                              "inline_payload": {
                                "block": {
                                  "attributes": {
                                    "description": {
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
                                  "block_types": {
                                    "input_schema": {
                                      "block": {
                                        "attributes": {
                                          "description": {
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
                                          "items": {
                                            "block": {
                                              "attributes": {
                                                "description": {
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
                                                "items": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
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
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "property": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "required": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
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
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "property": {
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
                                                },
                                                "required": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "type": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "items": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
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
                                                      "items": {
                                                        "block": {
                                                          "attributes": {
                                                            "description": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "items_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "properties_json": {
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
                                                          "description_kind": "plain"
                                                        },
                                                        "nesting_mode": "list"
                                                      },
                                                      "property": {
                                                        "block": {
                                                          "attributes": {
                                                            "description": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "items_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "name": {
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "properties_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "required": {
                                                              "computed": true,
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "bool"
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "property": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "required": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
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
                                    },
                                    "output_schema": {
                                      "block": {
                                        "attributes": {
                                          "description": {
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
                                          "items": {
                                            "block": {
                                              "attributes": {
                                                "description": {
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
                                                "items": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
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
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "property": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "required": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
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
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "property": {
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
                                                },
                                                "required": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "type": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "items": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
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
                                                      "items": {
                                                        "block": {
                                                          "attributes": {
                                                            "description": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "items_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "properties_json": {
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
                                                          "description_kind": "plain"
                                                        },
                                                        "nesting_mode": "list"
                                                      },
                                                      "property": {
                                                        "block": {
                                                          "attributes": {
                                                            "description": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "items_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "name": {
                                                              "description_kind": "plain",
                                                              "required": true,
                                                              "type": "string"
                                                            },
                                                            "properties_json": {
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "required": {
                                                              "computed": true,
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "bool"
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "property": {
                                                  "block": {
                                                    "attributes": {
                                                      "description": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "items_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "properties_json": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "required": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
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
                              "s3": {
                                "block": {
                                  "attributes": {
                                    "bucket_owner_account_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "uri": {
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
                  "open_api_schema": {
                    "block": {
                      "block_types": {
                        "inline_payload": {
                          "block": {
                            "attributes": {
                              "payload": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "s3": {
                          "block": {
                            "attributes": {
                              "bucket_owner_account_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "uri": {
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
                  "smithy_model": {
                    "block": {
                      "block_types": {
                        "inline_payload": {
                          "block": {
                            "attributes": {
                              "payload": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "s3": {
                          "block": {
                            "attributes": {
                              "bucket_owner_account_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "uri": {
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

func AwsBedrockagentcoreGatewayTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreGatewayTarget), &result)
	return &result
}
