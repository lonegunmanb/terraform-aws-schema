package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkfirewallFirewallPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "policy_variables": [
                "list",
                [
                  "object",
                  {
                    "rule_variables": [
                      "set",
                      [
                        "object",
                        {
                          "ip_set": [
                            "list",
                            [
                              "object",
                              {
                                "definition": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "key": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "stateful_default_actions": [
                "set",
                "string"
              ],
              "stateful_engine_options": [
                "list",
                [
                  "object",
                  {
                    "rule_order": "string",
                    "stream_exception_policy": "string"
                  }
                ]
              ],
              "stateful_rule_group_reference": [
                "set",
                [
                  "object",
                  {
                    "override": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string"
                        }
                      ]
                    ],
                    "priority": "number",
                    "resource_arn": "string"
                  }
                ]
              ],
              "stateless_custom_action": [
                "set",
                [
                  "object",
                  {
                    "action_definition": [
                      "list",
                      [
                        "object",
                        {
                          "publish_metric_action": [
                            "list",
                            [
                              "object",
                              {
                                "dimension": [
                                  "set",
                                  [
                                    "object",
                                    {
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "action_name": "string"
                  }
                ]
              ],
              "stateless_default_actions": [
                "set",
                "string"
              ],
              "stateless_fragment_default_actions": [
                "set",
                "string"
              ],
              "stateless_rule_group_reference": [
                "set",
                [
                  "object",
                  {
                    "priority": "number",
                    "resource_arn": "string"
                  }
                ]
              ],
              "tls_inspection_configuration_arn": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkfirewallFirewallPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkfirewallFirewallPolicy), &result)
	return &result
}
