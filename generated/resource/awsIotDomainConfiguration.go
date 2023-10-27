package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotDomainConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_type": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_certificate_arns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "service_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      "validation_certificate_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "authorizer_config": {
        "block": {
          "attributes": {
            "allow_authorizer_override": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "default_authorizer_name": {
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
      "tls_config": {
        "block": {
          "attributes": {
            "security_policy": {
              "computed": true,
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

func AwsIotDomainConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotDomainConfiguration), &result)
	return &result
}
