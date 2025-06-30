package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRolesanywhereTrustAnchor = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      }
    },
    "block_types": {
      "notification_settings": {
        "block": {
          "attributes": {
            "channel": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "configured_by": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "event": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 50,
        "nesting_mode": "set"
      },
      "source": {
        "block": {
          "attributes": {
            "source_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "source_data": {
              "block": {
                "attributes": {
                  "acm_pca_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "x509_certificate_data": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRolesanywhereTrustAnchorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRolesanywhereTrustAnchor), &result)
	return &result
}
