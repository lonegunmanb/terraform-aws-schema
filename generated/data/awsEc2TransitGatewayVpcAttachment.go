package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayVpcAttachment = `{
  "block": {
    "attributes": {
      "dns_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
                "list",
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

func AwsEc2TransitGatewayVpcAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayVpcAttachment), &result)
	return &result
}