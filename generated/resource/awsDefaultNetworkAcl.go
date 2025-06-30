package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDefaultNetworkAcl = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_network_acl_id": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "egress": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_port": {
              "description_kind": "plain",
              "required": true,
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
            "ipv6_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_no": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ingress": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_port": {
              "description_kind": "plain",
              "required": true,
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
            "ipv6_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_no": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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

func AwsDefaultNetworkAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDefaultNetworkAcl), &result)
	return &result
}
