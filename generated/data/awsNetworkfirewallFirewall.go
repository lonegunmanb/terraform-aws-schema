package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkfirewallFirewall = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "availability_zone_id": "string"
            }
          ]
        ]
      },
      "delete_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_analysis_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "key_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "firewall_policy_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "firewall_status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_usage_summary": [
                "set",
                [
                  "object",
                  {
                    "cidrs": [
                      "set",
                      [
                        "object",
                        {
                          "available_cidr_count": "number",
                          "ip_set_references": [
                            "set",
                            [
                              "object",
                              {
                                "resolved_cidr_count": "number"
                              }
                            ]
                          ],
                          "utilized_cidr_count": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "configuration_sync_state_summary": "string",
              "status": "string",
              "sync_states": [
                "set",
                [
                  "object",
                  {
                    "attachment": [
                      "list",
                      [
                        "object",
                        {
                          "endpoint_id": "string",
                          "status": "string",
                          "status_message": "string",
                          "subnet_id": "string"
                        }
                      ]
                    ],
                    "availability_zone": "string"
                  }
                ]
              ],
              "transit_gateway_attachment_sync_states": [
                "list",
                [
                  "object",
                  {
                    "attachment_id": "string",
                    "status_message": "string",
                    "transit_gateway_attachment_status": "string"
                  }
                ]
              ]
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
      },
      "subnet_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "subnet_id": "string"
            }
          ]
        ]
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
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkfirewallFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkfirewallFirewall), &result)
	return &result
}
