package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayPolicyTableEntry = `{
  "block": {
    "attributes": {
      "policy_rule_number": {
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
      "target_route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_policy_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "policy_rule": {
        "block": {
          "attributes": {
            "destination_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "metadata": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
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

func AwsEc2TransitGatewayPolicyTableEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayPolicyTableEntry), &result)
	return &result
}
