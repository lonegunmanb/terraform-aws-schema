package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDefaultSecurityGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
              "cidr_blocks": [
                "list",
                "string"
              ],
              "description": "string",
              "from_port": "number",
              "ipv6_cidr_blocks": [
                "list",
                "string"
              ],
              "prefix_list_ids": [
                "list",
                "string"
              ],
              "protocol": "string",
              "security_groups": [
                "set",
                "string"
              ],
              "self": "bool",
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
              "cidr_blocks": [
                "list",
                "string"
              ],
              "description": "string",
              "from_port": "number",
              "ipv6_cidr_blocks": [
                "list",
                "string"
              ],
              "prefix_list_ids": [
                "list",
                "string"
              ],
              "protocol": "string",
              "security_groups": [
                "set",
                "string"
              ],
              "self": "bool",
              "to_port": "number"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
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
      "revoke_rules_on_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsDefaultSecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDefaultSecurityGroup), &result)
	return &result
}
