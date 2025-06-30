package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConfigConfigRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "input_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_execution_frequency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
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
      "evaluation_mode": {
        "block": {
          "attributes": {
            "mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "scope": {
        "block": {
          "attributes": {
            "compliance_resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compliance_resource_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "tag_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value": {
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
      "source": {
        "block": {
          "attributes": {
            "owner": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_policy_details": {
              "block": {
                "attributes": {
                  "enable_debug_log_delivery": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "policy_runtime": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "policy_text": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "source_detail": {
              "block": {
                "attributes": {
                  "event_source": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "maximum_execution_frequency": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 25,
              "nesting_mode": "set"
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

func AwsConfigConfigRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConfigConfigRule), &result)
	return &result
}
