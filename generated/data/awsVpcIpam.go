package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpam = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_resource_discovery_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_resource_discovery_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_private_gua": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "operating_regions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "region_name": "string"
            }
          ]
        ]
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_default_scope_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_default_scope_id": {
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
      "resource_discovery_association_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scope_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
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
      "tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcIpamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcIpam), &result)
	return &result
}
