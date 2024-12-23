package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsoadminApplication = `{
  "block": {
    "attributes": {
      "application_account": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_provider_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "portal_options": {
        "block": {
          "attributes": {
            "visibility": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "sign_in_options": {
              "block": {
                "attributes": {
                  "application_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin": {
                    "computed": true,
                    "description_kind": "plain",
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
  "version": 0
}`

func AwsSsoadminApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsoadminApplication), &result)
	return &result
}
