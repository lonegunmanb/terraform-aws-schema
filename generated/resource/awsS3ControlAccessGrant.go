package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3ControlAccessGrant = `{
  "block": {
    "attributes": {
      "access_grant_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_grant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_grants_location_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "grant_scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "permission": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_prefix_type": {
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
      }
    },
    "block_types": {
      "access_grants_location_configuration": {
        "block": {
          "attributes": {
            "s3_sub_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "grantee": {
        "block": {
          "attributes": {
            "grantee_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "grantee_type": {
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
  "version": 0
}`

func AwsS3ControlAccessGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3ControlAccessGrant), &result)
	return &result
}
