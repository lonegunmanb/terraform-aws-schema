package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkAclRule = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "from_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "icmp_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "icmp_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_acl_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocol": {
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
      "rule_action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_number": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "to_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkAclRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkAclRule), &result)
	return &result
}
