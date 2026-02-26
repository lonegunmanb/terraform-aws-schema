package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2WebAclRuleGroupAssociation = `{
  "block": {
    "attributes": {
      "override_action": {
        "computed": true,
        "description": "Override action for the rule group. Valid values are 'none' and 'count'. Defaults to 'none'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "Priority of the rule within the Web ACL.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description": "Name of the rule to create in the Web ACL that references the rule group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "web_acl_arn": {
        "description": "ARN of the Web ACL to associate the Rule Group with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "managed_rule_group": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the managed rule group.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vendor_name": {
              "description": "Name of the managed rule group vendor.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description": "Version of the managed rule group. Omit this to use the default version.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "managed_rule_group_configs": {
              "block": {
                "block_types": {
                  "aws_managed_rules_acfp_rule_set": {
                    "block": {
                      "attributes": {
                        "creation_path": {
                          "description": "Path to the account creation endpoint on the protected website",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "enable_regex_in_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "registration_page_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "request_inspection": {
                          "block": {
                            "attributes": {
                              "payload_type": {
                                "description": "Payload type for inspection, either JSON or FORM_ENCODED.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "address_fields": {
                                "block": {
                                  "attributes": {
                                    "identifiers": {
                                      "description": "Identifiers of the address fields",
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
                              "email_field": {
                                "block": {
                                  "attributes": {
                                    "identifier": {
                                      "description": "Identifier of the email field",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "password_field": {
                                "block": {
                                  "attributes": {
                                    "identifier": {
                                      "description": "Identifier of the password field",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "phone_number_fields": {
                                "block": {
                                  "attributes": {
                                    "identifiers": {
                                      "description": "Identifiers of the phone number fields",
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
                              "username_field": {
                                "block": {
                                  "attributes": {
                                    "identifier": {
                                      "description": "Identifier of the username field",
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
                        "response_inspection": {
                          "block": {
                            "block_types": {
                              "body_contains": {
                                "block": {
                                  "attributes": {
                                    "failure_strings": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "success_strings": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "header": {
                                "block": {
                                  "attributes": {
                                    "failure_values": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "name": {
                                      "description": "Name of the HTTP header to inspect",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "success_values": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "json": {
                                "block": {
                                  "attributes": {
                                    "failure_values": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "identifier": {
                                      "description": "Identifier of the JSON field to inspect",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "success_values": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "status_code": {
                                "block": {
                                  "attributes": {
                                    "failure_codes": {
                                      "description": "Status codes that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "number"
                                      ]
                                    },
                                    "success_codes": {
                                      "description": "Status codes that indicate a successful login or account creation attempt",
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
                  "aws_managed_rules_anti_ddos_rule_set": {
                    "block": {
                      "attributes": {
                        "sensitivity_to_block": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "client_side_action_config": {
                          "block": {
                            "block_types": {
                              "challenge": {
                                "block": {
                                  "attributes": {
                                    "sensitivity": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "usage_of_action": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "exempt_uri_regular_expression": {
                                      "block": {
                                        "attributes": {
                                          "regex_string": {
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
                  "aws_managed_rules_atp_rule_set": {
                    "block": {
                      "attributes": {
                        "enable_regex_in_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "login_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "request_inspection": {
                          "block": {
                            "attributes": {
                              "payload_type": {
                                "description": "Payload type for inspection, either JSON or FORM_ENCODED.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "password_field": {
                                "block": {
                                  "attributes": {
                                    "identifier": {
                                      "description": "Identifier of the password field",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "username_field": {
                                "block": {
                                  "attributes": {
                                    "identifier": {
                                      "description": "Identifier of the username field",
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
                        "response_inspection": {
                          "block": {
                            "block_types": {
                              "body_contains": {
                                "block": {
                                  "attributes": {
                                    "failure_strings": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "success_strings": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "header": {
                                "block": {
                                  "attributes": {
                                    "failure_values": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "name": {
                                      "description": "Name of the HTTP header to inspect",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "success_values": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "json": {
                                "block": {
                                  "attributes": {
                                    "failure_values": {
                                      "description": "Strings that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "identifier": {
                                      "description": "Identifier of the JSON field to inspect",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "success_values": {
                                      "description": "Strings that indicate a successful login or account creation attempt",
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
                              "status_code": {
                                "block": {
                                  "attributes": {
                                    "failure_codes": {
                                      "description": "Status codes that indicate a failed login or account creation attempt",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "number"
                                      ]
                                    },
                                    "success_codes": {
                                      "description": "Status codes that indicate a successful login or account creation attempt",
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
                  "aws_managed_rules_bot_control_rule_set": {
                    "block": {
                      "attributes": {
                        "enable_machine_learning": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "inspection_level": {
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
            "rule_action_override": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the rule to override.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "action_to_use": {
                    "block": {
                      "block_types": {
                        "allow": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "block": {
                          "block": {
                            "block_types": {
                              "custom_response": {
                                "block": {
                                  "attributes": {
                                    "custom_response_body_key": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "response_code": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "response_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "captcha": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "challenge": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "count": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                      "description": "Action to use in place of the rule action.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Action settings to use in place of rule actions configured inside the rule group. You can specify up to 100 overrides.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Managed rule group configuration.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "rule_group_reference": {
        "block": {
          "attributes": {
            "arn": {
              "description": "ARN of the Rule Group to associate with the Web ACL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "rule_action_override": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the rule to override.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "action_to_use": {
                    "block": {
                      "block_types": {
                        "allow": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "block": {
                          "block": {
                            "block_types": {
                              "custom_response": {
                                "block": {
                                  "attributes": {
                                    "custom_response_body_key": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "response_code": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "response_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "captcha": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "challenge": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                        "count": {
                          "block": {
                            "block_types": {
                              "custom_request_handling": {
                                "block": {
                                  "block_types": {
                                    "insert_header": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
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
                      "description": "Action to use in place of the rule action.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Action settings to use in place of rule actions configured inside the rule group. You can specify up to 100 overrides.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Rule Group reference configuration.",
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
      "visibility_config": {
        "block": {
          "attributes": {
            "cloudwatch_metrics_enabled": {
              "description": "Indicates whether the rule is available for use in the metrics for the web ACL.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "metric_name": {
              "description": "A name for the metrics for this rule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sampled_requests_enabled": {
              "description": "Indicates whether to store a sampling of the web requests that match the rule.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Visibility configuration for the rule.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description": "Associates a WAFv2 Rule Group (custom or managed) with a Web ACL by adding a rule that references the Rule Group.",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWafv2WebAclRuleGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2WebAclRuleGroupAssociation), &result)
	return &result
}
