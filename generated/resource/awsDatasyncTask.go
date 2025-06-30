package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatasyncTask = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_log_group_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_location_arn": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_location_arn": {
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
      "task_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "excludes": {
        "block": {
          "attributes": {
            "filter_type": {
              "description_kind": "plain",
              "optional": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "includes": {
        "block": {
          "attributes": {
            "filter_type": {
              "description_kind": "plain",
              "optional": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "options": {
        "block": {
          "attributes": {
            "atime": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bytes_per_second": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gid": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mtime": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "object_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "overwrite_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "posix_permissions": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preserve_deleted_files": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preserve_devices": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_descriptor_copy_flags": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "task_queueing": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "transfer_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uid": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_mode": {
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
      "schedule": {
        "block": {
          "attributes": {
            "schedule_expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "task_report_config": {
        "block": {
          "attributes": {
            "output_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "report_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_object_versioning": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "report_overrides": {
              "block": {
                "attributes": {
                  "deleted_override": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "skipped_override": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transferred_override": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "verified_override": {
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
            "s3_destination": {
              "block": {
                "attributes": {
                  "bucket_access_role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subdirectory": {
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

func AwsDatasyncTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatasyncTask), &result)
	return &result
}
