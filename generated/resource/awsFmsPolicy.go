package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFmsPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_all_policy_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delete_unused_fm_managed_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_tags": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "remediation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_set_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "resource_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type_list": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      }
    },
    "block_types": {
      "exclude_map": {
        "block": {
          "attributes": {
            "account": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "orgunit": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "include_map": {
        "block": {
          "attributes": {
            "account": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "orgunit": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "security_service_policy_data": {
        "block": {
          "attributes": {
            "managed_service_data": {
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
            "policy_option": {
              "block": {
                "block_types": {
                  "network_acl_common_policy": {
                    "block": {
                      "block_types": {
                        "network_acl_entry_set": {
                          "block": {
                            "attributes": {
                              "force_remediate_for_first_entries": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              },
                              "force_remediate_for_last_entries": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "first_entry": {
                                "block": {
                                  "attributes": {
                                    "cidr_block": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "egress": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    },
                                    "ipv6_cidr_block": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "protocol": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "rule_action": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "icmp_type_code": {
                                      "block": {
                                        "attributes": {
                                          "code": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "port_range": {
                                      "block": {
                                        "attributes": {
                                          "from": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "to": {
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
                                "nesting_mode": "set"
                              },
                              "last_entry": {
                                "block": {
                                  "attributes": {
                                    "cidr_block": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "egress": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    },
                                    "ipv6_cidr_block": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "protocol": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "rule_action": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "icmp_type_code": {
                                      "block": {
                                        "attributes": {
                                          "code": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "port_range": {
                                      "block": {
                                        "attributes": {
                                          "from": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "to": {
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
                                "nesting_mode": "set"
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "network_firewall_policy": {
                    "block": {
                      "attributes": {
                        "firewall_deployment_model": {
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
                  "third_party_firewall_policy": {
                    "block": {
                      "attributes": {
                        "firewall_deployment_model": {
                          "description_kind": "plain",
                          "optional": true,
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsFmsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFmsPolicy), &result)
	return &result
}
