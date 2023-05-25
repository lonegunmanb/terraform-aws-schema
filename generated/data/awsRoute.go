package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute = `{
  "block": {
    "attributes": {
      "destination_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_ipv6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress_only_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_peering_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute), &result)
	return &result
}
