package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpnConnection = `{
  "block": {
    "attributes": {
      "category": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_attachment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_gateway_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_association_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pre_shared_key_arn": {
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
      "routes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "destination_cidr_block": "string",
              "source": "string",
              "state": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
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
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vgw_telemetries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accepted_route_count": "number",
              "last_status_change": "string",
              "outside_ip_address": "string",
              "status": "string",
              "status_message": "string"
            }
          ]
        ]
      },
      "vpn_concentrator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_gateway_id": {
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

func AwsVpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpnConnection), &result)
	return &result
}
