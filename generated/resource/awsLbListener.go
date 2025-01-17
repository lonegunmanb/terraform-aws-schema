package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbListener = `{
  "block": {
    "attributes": {
      "alpn_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_arn": {
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
      "load_balancer_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_issuer_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_leaf_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_serial_number_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_subject_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_mtls_clientcert_validity_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_tls_cipher_suite_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_request_x_amzn_tls_version_header_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_allow_credentials_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_allow_headers_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_allow_methods_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_allow_origin_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_expose_headers_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_access_control_max_age_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_content_security_policy_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_server_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "routing_http_response_strict_transport_security_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_x_content_type_options_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_http_response_x_frame_options_header_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_policy": {
        "computed": true,
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
      "tcp_idle_timeout_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "default_action": {
        "block": {
          "attributes": {
            "order": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_group_arn": {
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
            "authenticate_cognito": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "user_pool_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_pool_client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_pool_domain": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "authenticate_oidc": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "authorization_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "issuer": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "token_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_info_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "fixed_response": {
              "block": {
                "attributes": {
                  "content_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "message_body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
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
            "forward": {
              "block": {
                "block_types": {
                  "stickiness": {
                    "block": {
                      "attributes": {
                        "duration": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
                  },
                  "target_group": {
                    "block": {
                      "attributes": {
                        "arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "weight": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "redirect": {
              "block": {
                "attributes": {
                  "host": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "mutual_authentication": {
        "block": {
          "attributes": {
            "advertise_trust_store_ca_names": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_client_certificate_expiry": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "trust_store_arn": {
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
            "update": {
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

func AwsLbListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbListener), &result)
	return &result
}
