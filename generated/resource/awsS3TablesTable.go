package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3TablesTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "kms_key_arn": "string",
            "sse_algorithm": "string"
          }
        ]
      },
      "format": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "iceberg_compaction": [
              "object",
              {
                "settings": [
                  "object",
                  {
                    "target_file_size_mb": "number"
                  }
                ],
                "status": "string"
              }
            ],
            "iceberg_snapshot_management": [
              "object",
              {
                "settings": [
                  "object",
                  {
                    "max_snapshot_age_hours": "number",
                    "min_snapshots_to_keep": "number"
                  }
                ],
                "status": "string"
              }
            ]
          }
        ]
      },
      "metadata_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modified_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_bucket_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "warehouse_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "block_types": {
            "iceberg": {
              "block": {
                "block_types": {
                  "schema": {
                    "block": {
                      "block_types": {
                        "field": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The name of the field.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "A Boolean value that specifies whether values are required for each row in this field. Default: false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "type": {
                                "description": "The field type. S3 Tables supports all Apache Iceberg primitive types.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "List of schema fields for the Iceberg table.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Schema configuration for the Iceberg table.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Iceberg metadata configuration.",
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

func AwsS3TablesTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3TablesTable), &result)
	return &result
}
