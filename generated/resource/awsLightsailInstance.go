package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLightsailInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "blueprint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
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
      "ipv6_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_static_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ram_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLightsailInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLightsailInstance), &result)
	return &result
}