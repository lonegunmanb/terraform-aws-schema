package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcSecurityGroupRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_ipv4": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_ipv6": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "from_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_egress": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "prefix_list_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "referenced_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "to_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcSecurityGroupRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcSecurityGroupRule), &result)
	return &result
}
