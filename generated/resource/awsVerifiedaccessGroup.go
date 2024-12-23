package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedaccessGroup = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
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
      "verifiedaccess_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "verifiedaccess_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "verifiedaccess_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "sse_configuration": {
        "block": {
          "attributes": {
            "customer_managed_key_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
  "version": 0
}`

func AwsVerifiedaccessGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedaccessGroup), &result)
	return &result
}
