package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2CapacityBlockReservation = `{
  "block": {
    "attributes": {
      "arn": {
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
      "available_instance_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "capacity_block_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "commitment_info": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "commitment_end_date": "string",
            "committed_instance_count": "number"
          }
        ]
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_preference": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_date_type": {
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
      "instance_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_match_criteria": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "interruptible_capacity_allocation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "instance_count": "number",
            "interruptible_capacity_reservation_id": "string",
            "interruption_type": "string",
            "status": "string",
            "target_instance_count": "number"
          }
        ]
      },
      "interruption_info": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "interruption_type": "string",
            "source_capacity_reservation_id": "string"
          }
        ]
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
      "placement_group_arn": {
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
      "reservation_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description_kind": "plain",
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
        "type": [
          "map",
          "string"
        ]
      },
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
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

func AwsEc2CapacityBlockReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2CapacityBlockReservation), &result)
	return &result
}
