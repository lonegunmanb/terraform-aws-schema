package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmcontactsPlan = `{
  "block": {
    "attributes": {
      "contact_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
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
      }
    },
    "block_types": {
      "stage": {
        "block": {
          "attributes": {
            "duration_in_minutes": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "target": {
              "block": {
                "block_types": {
                  "channel_target_info": {
                    "block": {
                      "attributes": {
                        "contact_channel_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "retry_interval_in_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "contact_target_info": {
                    "block": {
                      "attributes": {
                        "contact_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_essential": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmcontactsPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmcontactsPlan), &result)
	return &result
}
