package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkAcl = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "egress": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "action": "string",
              "cidr_block": "string",
              "from_port": "number",
              "icmp_code": "number",
              "icmp_type": "number",
              "ipv6_cidr_block": "string",
              "protocol": "string",
              "rule_no": "number",
              "to_port": "number"
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
      "ingress": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "action": "string",
              "cidr_block": "string",
              "from_port": "number",
              "icmp_code": "number",
              "icmp_type": "number",
              "ipv6_cidr_block": "string",
              "protocol": "string",
              "rule_no": "number",
              "to_port": "number"
            }
          ]
        ]
      },
      "owner_id": {
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
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkAcl), &result)
	return &result
}
