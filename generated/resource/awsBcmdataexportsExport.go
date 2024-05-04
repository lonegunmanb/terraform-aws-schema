package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBcmdataexportsExport = `{
  "block": {
    "attributes": {
      "id": {
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
      }
    },
    "block_types": {
      "export": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_query": {
              "block": {
                "attributes": {
                  "query_statement": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_configurations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      [
                        "map",
                        "string"
                      ]
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "destination_configurations": {
              "block": {
                "block_types": {
                  "s3_destination": {
                    "block": {
                      "attributes": {
                        "s3_bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_prefix": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_region": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "s3_output_configurations": {
                          "block": {
                            "attributes": {
                              "compression": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "format": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "output_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "overwrite": {
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
            "refresh_cadence": {
              "block": {
                "attributes": {
                  "frequency": {
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsBcmdataexportsExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBcmdataexportsExport), &result)
	return &result
}
