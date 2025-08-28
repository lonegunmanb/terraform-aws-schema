package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpams = `{
  "block": {
    "attributes": {
      "ipam_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ipams": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "default_resource_discovery_association_id": "string",
              "default_resource_discovery_id": "string",
              "description": "string",
              "enable_private_gua": "bool",
              "id": "string",
              "ipam_region": "string",
              "metered_account": "string",
              "operating_regions": [
                "list",
                [
                  "object",
                  {
                    "region_name": "string"
                  }
                ]
              ],
              "owner_id": "string",
              "private_default_scope_id": "string",
              "public_default_scope_id": "string",
              "resource_discovery_association_count": "number",
              "scope_count": "number",
              "state": "string",
              "state_message": "string",
              "tier": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcIpamsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcIpams), &result)
	return &result
}
