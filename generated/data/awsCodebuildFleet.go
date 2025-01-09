package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodebuildFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "base_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "compute_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk": "number",
              "machine_type": "string",
              "memory": "number",
              "vcpu": "number"
            }
          ]
        ]
      },
      "compute_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_service_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "overflow_behavior": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "desired_capacity": "number",
              "max_capacity": "number",
              "scaling_type": "string",
              "target_tracking_scaling_configs": [
                "list",
                [
                  "object",
                  {
                    "metric_type": "string",
                    "target_value": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "context": "string",
              "message": "string",
              "status_code": "string"
            }
          ]
        ]
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
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "security_group_ids": [
                "set",
                "string"
              ],
              "subnets": [
                "set",
                "string"
              ],
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodebuildFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodebuildFleet), &result)
	return &result
}
