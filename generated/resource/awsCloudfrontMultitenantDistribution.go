package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontMultitenantDistribution = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "caller_reference": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_root_object": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "in_progress_invalidation_batches": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
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
      "web_acl_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "active_trusted_key_groups": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "block_types": {
            "items": {
              "block": {
                "attributes": {
                  "key_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_pair_ids": {
                    "computed": true,
                    "description_kind": "plain",
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
      },
      "cache_behavior": {
        "block": {
          "attributes": {
            "cache_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compress": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "field_level_encryption_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "origin_request_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "realtime_log_config_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_headers_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "viewer_protocol_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "allowed_methods": {
              "block": {
                "attributes": {
                  "cached_methods": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "items": {
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
            "function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "lambda_function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "include_body": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "lambda_function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trusted_key_groups": {
              "block": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "items": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "custom_error_response": {
        "block": {
          "attributes": {
            "error_caching_min_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "error_code": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "response_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_page_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "default_cache_behavior": {
        "block": {
          "attributes": {
            "cache_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compress": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "field_level_encryption_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "origin_request_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "realtime_log_config_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_headers_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "viewer_protocol_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "allowed_methods": {
              "block": {
                "attributes": {
                  "cached_methods": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "items": {
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
            "function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "lambda_function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "include_body": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "lambda_function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trusted_key_groups": {
              "block": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "items": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "origin": {
        "block": {
          "attributes": {
            "connection_attempts": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "connection_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_access_control_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "origin_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_completion_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "custom_header": {
              "block": {
                "attributes": {
                  "header_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "header_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "custom_origin_config": {
              "block": {
                "attributes": {
                  "http_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "https_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "ip_address_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "origin_keepalive_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_protocol_policy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "origin_read_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_ssl_protocols": {
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
            "origin_shield": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "origin_shield_region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vpc_origin_config": {
              "block": {
                "attributes": {
                  "origin_keepalive_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_read_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "vpc_origin_id": {
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
      "origin_group": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "failover_criteria": {
              "block": {
                "attributes": {
                  "status_codes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "member": {
              "block": {
                "attributes": {
                  "origin_id": {
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
      "restrictions": {
        "block": {
          "block_types": {
            "geo_restriction": {
              "block": {
                "attributes": {
                  "items": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "restriction_type": {
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
      "tenant_config": {
        "block": {
          "block_types": {
            "parameter_definition": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "definition": {
                    "block": {
                      "block_types": {
                        "string_schema": {
                          "block": {
                            "attributes": {
                              "comment": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "default_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "description_kind": "plain",
                                "required": true,
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
      "viewer_certificate": {
        "block": {
          "attributes": {
            "acm_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cloudfront_default_certificate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "minimum_protocol_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_support_method": {
              "computed": true,
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
  "version": 0
}`

func AwsCloudfrontMultitenantDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontMultitenantDistribution), &result)
	return &result
}
