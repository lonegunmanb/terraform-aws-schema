package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNatGateway = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_provision_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "allocation_ids": [
                "set",
                "string"
              ],
              "availability_zone": "string",
              "availability_zone_id": "string"
            }
          ]
        ]
      },
      "connectivity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
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
      "regional_nat_gateway_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "allocation_id": "string",
              "association_id": "string",
              "availability_zone": "string",
              "availability_zone_id": "string",
              "network_interface_id": "string",
              "public_ip": "string",
              "status": "string"
            }
          ]
        ]
      },
      "route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_allocation_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "secondary_private_ip_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "secondary_private_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      },
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNatGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNatGateway), &result)
	return &result
}
