package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpc = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assign_generated_ipv6_cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_network_acl_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dhcp_options_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_dns_hostnames": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_dns_support": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_network_address_usage_metrics": {
        "computed": true,
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
      "instance_tenancy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv4_ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv4_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv6_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block_network_border_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "main_route_table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
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
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsVpcSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpc), &result)
	return &result
}
