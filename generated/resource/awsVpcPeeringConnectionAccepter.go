package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcPeeringConnectionAccepter = `{
  "block": {
    "attributes": {
      "accept_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_accept": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "peer_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "peer_vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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
      },
      "vpc_peering_connection_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "accepter": {
        "block": {
          "attributes": {
            "allow_classic_link_to_remote_vpc": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_remote_vpc_dns_resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_vpc_to_remote_classic_link": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "requester": {
        "block": {
          "attributes": {
            "allow_classic_link_to_remote_vpc": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_remote_vpc_dns_resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_vpc_to_remote_classic_link": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcPeeringConnectionAccepterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcPeeringConnectionAccepter), &result)
	return &result
}