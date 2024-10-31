package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbListenerRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "action": {
        "block": {
          "attributes": {
            "order": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "authenticate_cognito": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "user_pool_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_pool_client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_pool_domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "authenticate_oidc": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "authorization_endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "issuer": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "token_endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_info_endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "fixed_response": {
              "block": {
                "attributes": {
                  "content_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "message_body": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "forward": {
              "block": {
                "block_types": {
                  "stickiness": {
                    "block": {
                      "attributes": {
                        "duration": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "single"
                  },
                  "target_group": {
                    "block": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "redirect": {
              "block": {
                "attributes": {
                  "host": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "query": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status_code": {
                    "computed": true,
                    "description_kind": "plain",
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
        "nesting_mode": "list"
      },
      "condition": {
        "block": {
          "block_types": {
            "host_header": {
              "block": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "http_header": {
              "block": {
                "attributes": {
                  "http_header_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "http_request_method": {
              "block": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "path_pattern": {
              "block": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "query_string": {
              "block": {
                "block_types": {
                  "values": {
                    "block": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
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
              "nesting_mode": "single"
            },
            "source_ip": {
              "block": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLbListenerRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbListenerRule), &result)
	return &result
}
