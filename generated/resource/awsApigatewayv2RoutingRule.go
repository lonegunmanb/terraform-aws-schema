package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2RoutingRule = `{
  "block": {
    "attributes": {
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_rule_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "routing_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "action": {
        "block": {
          "block_types": {
            "invoke_api": {
              "block": {
                "attributes": {
                  "api_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stage": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "strip_base_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      },
      "condition": {
        "block": {
          "block_types": {
            "match_base_paths": {
              "block": {
                "attributes": {
                  "any_of": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "match_headers": {
              "block": {
                "block_types": {
                  "any_of": {
                    "block": {
                      "attributes": {
                        "header": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_glob": {
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

func AwsApigatewayv2RoutingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2RoutingRule), &result)
	return &result
}
