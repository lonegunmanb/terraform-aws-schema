package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotCaCertificate = `{
  "block": {
    "attributes": {
      "active": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "allow_auto_registration": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ca_certificate_pem": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "certificate_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "customer_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "generation_id": {
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
      "validity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "not_after": "string",
              "not_before": "string"
            }
          ]
        ]
      },
      "verification_certificate_pem": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      }
    },
    "block_types": {
      "registration_config": {
        "block": {
          "attributes": {
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "template_body": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "template_name": {
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

func AwsIotCaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotCaCertificate), &result)
	return &result
}
