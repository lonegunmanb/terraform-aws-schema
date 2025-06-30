package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEip = `{
  "block": {
    "attributes": {
      "address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "allocation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associate_with_private_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "carrier_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ipv4_pool": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
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
      "instance": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_border_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ptr_record": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ipv4_pool": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
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
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsEipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEip), &result)
	return &result
}
