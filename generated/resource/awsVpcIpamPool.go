package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpamPool = `{
  "block": {
    "attributes": {
      "address_family": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allocation_default_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_max_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_min_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_resource_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_import": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "aws_service": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cascade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
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
      "ipam_scope_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_scope_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_depth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "public_ip_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_advertisable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
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
    "block_types": {
      "source_resource": {
        "block": {
          "attributes": {
            "resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_owner": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AwsVpcIpamPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcIpamPool), &result)
	return &result
}
