package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsMigrationProject = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_profile_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
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
      "transformation_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "schema_conversion_application_attributes": {
        "block": {
          "attributes": {
            "s3_bucket_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_bucket_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "source_data_provider_descriptor": {
        "block": {
          "attributes": {
            "data_provider_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "data_provider_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "target_data_provider_descriptor": {
        "block": {
          "attributes": {
            "data_provider_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "data_provider_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AwsDmsMigrationProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsMigrationProject), &result)
	return &result
}
