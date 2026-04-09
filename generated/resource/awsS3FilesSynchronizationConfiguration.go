package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3FilesSynchronizationConfiguration = `{
  "block": {
    "attributes": {
      "file_system_id": {
        "description": "File system ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "latest_version_number": {
        "computed": true,
        "description": "Latest version number for optimistic locking",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "expiration_data_rule": {
        "block": {
          "attributes": {
            "days_after_last_access": {
              "description": "Days after last access before data expires",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "import_data_rule": {
        "block": {
          "attributes": {
            "prefix": {
              "description": "S3 prefix for import",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "size_less_than": {
              "description": "Maximum file size to import",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "trigger": {
              "description": "Import trigger type",
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
  "version": 0
}`

func AwsS3FilesSynchronizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3FilesSynchronizationConfiguration), &result)
	return &result
}
