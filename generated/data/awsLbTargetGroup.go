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
      "connection_termination": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "deregistration_delay": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "load_balancer_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "load_balancing_algorithm_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancing_anomaly_mitigation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_cross_zone_enabled": {
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
      "preserve_client_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "proxy_protocol_v2": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
              "cookie_name": "string",
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
    "block_types": {
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

func AwsLbTargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbTargetGroup), &result)
	return &result
}
