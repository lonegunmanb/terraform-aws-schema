package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmanagerCoreNetworkPolicyDocument = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "json": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "attachment_policies": {
        "block": {
          "attributes": {
            "condition_logic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_number": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "add_to_network_function_group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "association_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "require_acceptance": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "segment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_value_of_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "conditions": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "core_network_configuration": {
        "block": {
          "attributes": {
            "asn_ranges": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "dns_support": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "inside_cidr_blocks": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "security_group_referencing_support": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vpn_ecmp_support": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "edge_locations": {
              "block": {
                "attributes": {
                  "asn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "inside_cidr_blocks": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "location": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "network_function_groups": {
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
            "require_attachment_acceptance": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "segment_actions": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_cidr_blocks": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "destinations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "segment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "share_with": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "share_with_except": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "via": {
              "block": {
                "attributes": {
                  "network_function_groups": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "with_edge_override": {
                    "block": {
                      "attributes": {
                        "edge_sets": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            [
                              "set",
                              "string"
                            ]
                          ]
                        },
                        "use_edge": {
                          "deprecated": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "use_edge_location": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "when_sent_to": {
              "block": {
                "attributes": {
                  "segments": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "segments": {
        "block": {
          "attributes": {
            "allow_filter": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "deny_filter": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "edge_locations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "isolate_attachments": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "require_attachment_acceptance": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkmanagerCoreNetworkPolicyDocumentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmanagerCoreNetworkPolicyDocument), &result)
	return &result
}
