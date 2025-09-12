package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueCatalogTableOptimizer = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_name": {
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
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "orphan_file_deletion_configuration": {
              "block": {
                "block_types": {
                  "iceberg_configuration": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "orphan_file_retention_period_in_days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "run_rate_in_hours": {
                          "computed": true,
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
              "nesting_mode": "list"
            },
            "retention_configuration": {
              "block": {
                "block_types": {
                  "iceberg_configuration": {
                    "block": {
                      "attributes": {
                        "clean_expired_files": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "number_of_snapshots_to_retain": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "run_rate_in_hours": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "snapshot_retention_period_in_days": {
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

func AwsGlueCatalogTableOptimizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueCatalogTableOptimizer), &result)
	return &result
}
