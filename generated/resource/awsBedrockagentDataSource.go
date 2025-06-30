package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentDataSource = `{
  "block": {
    "attributes": {
      "data_deletion_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "knowledge_base_id": {
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
      }
    },
    "block_types": {
      "data_source_configuration": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "confluence_configuration": {
              "block": {
                "block_types": {
                  "crawler_configuration": {
                    "block": {
                      "block_types": {
                        "filter_configuration": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "pattern_object_filter": {
                                "block": {
                                  "block_types": {
                                    "filters": {
                                      "block": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
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
                  "source_configuration": {
                    "block": {
                      "attributes": {
                        "auth_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "host_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "host_url": {
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
            "s3_configuration": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "bucket_owner_account_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "inclusion_prefixes": {
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
            },
            "salesforce_configuration": {
              "block": {
                "block_types": {
                  "crawler_configuration": {
                    "block": {
                      "block_types": {
                        "filter_configuration": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "pattern_object_filter": {
                                "block": {
                                  "block_types": {
                                    "filters": {
                                      "block": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
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
                  "source_configuration": {
                    "block": {
                      "attributes": {
                        "auth_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "host_url": {
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
            "share_point_configuration": {
              "block": {
                "block_types": {
                  "crawler_configuration": {
                    "block": {
                      "block_types": {
                        "filter_configuration": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "pattern_object_filter": {
                                "block": {
                                  "block_types": {
                                    "filters": {
                                      "block": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
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
                  "source_configuration": {
                    "block": {
                      "attributes": {
                        "auth_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "domain": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "host_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "site_urls": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "tenant_id": {
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
            "web_configuration": {
              "block": {
                "block_types": {
                  "crawler_configuration": {
                    "block": {
                      "attributes": {
                        "exclusion_filters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "inclusion_filters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "scope": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_agent": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "crawler_limits": {
                          "block": {
                            "attributes": {
                              "max_pages": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "rate_limit": {
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
                  "source_configuration": {
                    "block": {
                      "block_types": {
                        "url_configuration": {
                          "block": {
                            "block_types": {
                              "seed_urls": {
                                "block": {
                                  "attributes": {
                                    "url": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "server_side_encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "vector_ingestion_configuration": {
        "block": {
          "block_types": {
            "chunking_configuration": {
              "block": {
                "attributes": {
                  "chunking_strategy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "fixed_size_chunking_configuration": {
                    "block": {
                      "attributes": {
                        "max_tokens": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "overlap_percentage": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "hierarchical_chunking_configuration": {
                    "block": {
                      "attributes": {
                        "overlap_tokens": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "level_configuration": {
                          "block": {
                            "attributes": {
                              "max_tokens": {
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
                  },
                  "semantic_chunking_configuration": {
                    "block": {
                      "attributes": {
                        "breakpoint_percentile_threshold": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "buffer_size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "max_token": {
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
            },
            "custom_transformation_configuration": {
              "block": {
                "block_types": {
                  "intermediate_storage": {
                    "block": {
                      "block_types": {
                        "s3_location": {
                          "block": {
                            "attributes": {
                              "uri": {
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
                  "transformation": {
                    "block": {
                      "attributes": {
                        "step_to_apply": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "transformation_function": {
                          "block": {
                            "block_types": {
                              "transformation_lambda_configuration": {
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
            "parsing_configuration": {
              "block": {
                "attributes": {
                  "parsing_strategy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "bedrock_foundation_model_configuration": {
                    "block": {
                      "attributes": {
                        "model_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parsing_prompt": {
                          "block": {
                            "attributes": {
                              "parsing_prompt_string": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockagentDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentDataSource), &result)
	return &result
}
