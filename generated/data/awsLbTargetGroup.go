package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbTargetGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn_suffix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deregistration_delay": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "health_check": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "healthy_threshold": "number",
              "interval": "number",
              "matcher": "string",
              "path": "string",
              "port": "string",
              "protocol": "string",
              "timeout": "number",
              "unhealthy_threshold": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_multi_value_headers_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "load_balancing_algorithm_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "proxy_protocol_v2": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "slow_start": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "stickiness": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cookie_duration": "number",
              "enabled": "bool",
              "type": "string"
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
      "target_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLbTargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbTargetGroup), &result)
	return &result
}