package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketMetadataConfiguration = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expected_bucket_owner": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "metadata_configuration": {
        "block": {
          "attributes": {
            "destination": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "table_bucket_arn": "string",
                    "table_bucket_type": "string",
                    "table_namespace": "string"
                  }
                ]
              ]
            }
          },
          "block_types": {
            "inventory_table_configuration": {
              "block": {
                "attributes": {
                  "configuration_state": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption_configuration": {
                    "block": {
                      "attributes": {
                        "kms_key_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sse_algorithm": {
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
            "journal_table_configuration": {
              "block": {
                "attributes": {
                  "table_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "encryption_configuration": {
                    "block": {
                      "attributes": {
                        "kms_key_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sse_algorithm": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "record_expiration": {
                    "block": {
                      "attributes": {
                        "days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "expiration": {
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AwsS3BucketMetadataConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketMetadataConfiguration), &result)
	return &result
}
