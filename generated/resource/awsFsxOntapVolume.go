package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxOntapVolume = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bypass_snaplock_enterprise_retention": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "copy_tags_to_backups": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "flexcache_endpoint_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "junction_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ontap_volume_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_style": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size_in_megabytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "skip_final_backup": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_efficiency_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_virtual_machine_id": {
        "description_kind": "plain",
        "required": true,
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
      "uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_style": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "aggregate_configuration": {
        "block": {
          "attributes": {
            "aggregates": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "constituents_per_aggregate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total_constituents": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "snaplock_configuration": {
        "block": {
          "attributes": {
            "audit_log_volume": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "privileged_delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "snaplock_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "volume_append_mode_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "autocommit_period": {
              "block": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "retention_period": {
              "block": {
                "block_types": {
                  "default_retention": {
                    "block": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "maximum_retention": {
                    "block": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "minimum_retention": {
                    "block": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
        "nesting_mode": "list"
      },
      "tiering_policy": {
        "block": {
          "attributes": {
            "cooling_period": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsFsxOntapVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxOntapVolume), &result)
	return &result
}
