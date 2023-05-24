package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSubnet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assign_ipv6_address_on_creation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "customer_owned_ipv4_pool": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_dns64": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_lni_at_device_index": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_resource_name_dns_a_record_on_launch": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_resource_name_dns_aaaa_record_on_launch": {
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
      "ipv6_cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_native": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "map_customer_owned_ip_on_launch": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "map_public_ip_on_launch": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "outpost_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_hostname_type_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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
  "version": 1
}`

func AwsSubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSubnet), &result)
	return &result
}
