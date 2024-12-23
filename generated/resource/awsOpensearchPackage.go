package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchPackage = `{
  "block": {
    "attributes": {
      "available_package_version": {
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
      "package_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "package_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "package_source": {
        "block": {
          "attributes": {
            "s3_bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AwsOpensearchPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchPackage), &result)
	return &result
}
