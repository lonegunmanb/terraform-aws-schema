package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkfirewallTlsInspectionConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_arn": "string",
              "certificate_serial": "string",
              "status": "string",
              "status_message": "string"
            }
          ]
        ]
      },
      "certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_arn": "string",
              "certificate_serial": "string",
              "status": "string",
              "status_message": "string"
            }
          ]
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "key_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_associations": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "tls_inspection_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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
      "tls_inspection_configuration": {
        "block": {
          "block_types": {
            "server_certificate_configuration": {
              "block": {
                "attributes": {
                  "certificate_authority_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "check_certificate_revocation_status": {
                    "block": {
                      "attributes": {
                        "revoked_status_action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "unknown_status_action": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "scope": {
                    "block": {
                      "attributes": {
                        "protocols": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        }
                      },
                      "block_types": {
                        "destination": {
                          "block": {
                            "attributes": {
                              "address_definition": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "destination_ports": {
                          "block": {
                            "attributes": {
                              "from_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "to_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "source": {
                          "block": {
                            "attributes": {
                              "address_definition": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "source_ports": {
                          "block": {
                            "attributes": {
                              "from_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "to_port": {
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
                  "server_certificate": {
                    "block": {
                      "attributes": {
                        "resource_arn": {
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
  "version": 0
}`

func AwsNetworkfirewallTlsInspectionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkfirewallTlsInspectionConfiguration), &result)
	return &result
}
