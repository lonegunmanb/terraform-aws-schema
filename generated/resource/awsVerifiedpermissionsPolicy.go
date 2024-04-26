package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedpermissionsPolicy = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "definition": {
        "block": {
          "block_types": {
            "static": {
              "block": {
                "attributes": {
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statement": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "template_linked": {
              "block": {
                "attributes": {
                  "policy_template_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "principal": {
                    "block": {
                      "attributes": {
                        "entity_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "entity_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "resource": {
                    "block": {
                      "attributes": {
                        "entity_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "entity_type": {
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

func AwsVerifiedpermissionsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedpermissionsPolicy), &result)
	return &result
}
