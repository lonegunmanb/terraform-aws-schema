package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2Host = `{
  "block": {
    "attributes": {
      "allocation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "allows_multiple_instance_types": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "asset_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_placement": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "available_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_instance_capacity": [
                "list",
                [
                  "object",
                  {
                    "available_capacity": "number",
                    "instance_type": "string",
                    "total_capacity": "number"
                  }
                ]
              ],
              "available_vcpus": "number"
            }
          ]
        ]
      },
      "cores": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "host_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_maintenance": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_recovery": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_reservation_id": {
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
      "instance_family": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instance_id": "string",
              "instance_type": "string",
              "owner_id": "string"
            }
          ]
        ]
      },
      "member_of_service_linked_resource_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
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
      "release_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sockets": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "total_vcpus": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      },
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AwsEc2HostSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2Host), &result)
	return &result
}
