package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMailmanagerRuleSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modification_date": {
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
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "block_types": {
                  "add_header": {
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
                  "archive": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "target_archive": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "bounce": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "diagnostic_message": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "sender": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "smtp_reply_code": {
                          "description_kind": "plain",
                          "required": true,
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
                    "nesting_mode": "list"
                  },
                  "deliver_to_mailbox": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mailbox_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "deliver_to_q_business": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "application_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "index_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "drop": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "invoke_lambda": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "function_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "invocation_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "retry_time_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "publish_to_sns": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "encoding": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "payload_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
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
                  "relay": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mail_from": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "relay": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "replace_recipient": {
                    "block": {
                      "attributes": {
                        "replace_with": {
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
                  },
                  "send": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "write_to_s3": {
                    "block": {
                      "attributes": {
                        "action_failure_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_sse_kms_key_id": {
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
            "condition": {
              "block": {
                "block_types": {
                  "boolean_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "is_in_address_list": {
                                "block": {
                                  "attributes": {
                                    "address_lists": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "attribute": {
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
                  "dmarc_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                  },
                  "ip_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
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
                  "number_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
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
                  "string_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_certificate_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "mime_header_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
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
                  "verdict_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
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
            "unless": {
              "block": {
                "block_types": {
                  "boolean_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "is_in_address_list": {
                                "block": {
                                  "attributes": {
                                    "address_lists": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "attribute": {
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
                  "dmarc_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                  },
                  "ip_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
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
                  "number_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
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
                  "string_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_certificate_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "mime_header_attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
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
                  "verdict_expression": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "evaluate": {
                          "block": {
                            "attributes": {
                              "attribute": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "analysis": {
                                "block": {
                                  "attributes": {
                                    "analyzer": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "result_field": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMailmanagerRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMailmanagerRuleSet), &result)
	return &result
}
