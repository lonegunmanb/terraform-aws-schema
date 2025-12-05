package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBackupPlan = `{
  "block": {
    "attributes": {
      "arn": {
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "advanced_backup_setting": {
        "block": {
          "attributes": {
            "backup_options": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "resource_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rule": {
        "block": {
          "attributes": {
            "completion_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_continuous_backup": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "recovery_point_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "rule_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression_timezone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_logically_air_gapped_backup_vault_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_vault_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "copy_action": {
              "block": {
                "attributes": {
                  "destination_vault_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "lifecycle": {
                    "block": {
                      "attributes": {
                        "cold_storage_after": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "opt_in_to_archive_for_supported_resources": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
              "nesting_mode": "set"
            },
            "lifecycle": {
              "block": {
                "attributes": {
                  "cold_storage_after": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "delete_after": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "opt_in_to_archive_for_supported_resources": {
                    "computed": true,
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
            "scan_action": {
              "block": {
                "attributes": {
                  "malware_scanner": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "scan_mode": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "set"
      },
      "scan_setting": {
        "block": {
          "attributes": {
            "malware_scanner": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "scanner_role_arn": {
              "description_kind": "plain",
              "required": true,
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
  "version": 1
}`

func AwsBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBackupPlan), &result)
	return &result
}
