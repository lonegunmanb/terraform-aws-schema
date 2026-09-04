package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDxConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_device": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bandwidth": {
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
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "partner_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_pool_size_ipv4": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "prefix_pool_size_ipv6": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "prefix_pool_unallocated_count_ipv4": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "prefix_pool_unallocated_count_ipv6": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "provider_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rate_limiter_status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "in_use": "number",
              "max_allowed": "number",
              "remaining": "number",
              "total_bandwidth": "string"
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
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
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
      "vlan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDxConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDxConnection), &result)
	return &result
}
