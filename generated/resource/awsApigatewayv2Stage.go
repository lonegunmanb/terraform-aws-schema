package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Stage = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_deploy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_certificate_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_arn": {
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
      "invoke_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
      "access_log_settings": {
        "block": {
          "attributes": {
            "destination_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "format": {
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
      "default_route_settings": {
        "block": {
          "attributes": {
            "data_trace_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "detailed_metrics_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "logging_level": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "throttling_burst_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throttling_rate_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "route_settings": {
        "block": {
          "attributes": {
            "data_trace_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "detailed_metrics_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "logging_level": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "route_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "throttling_burst_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throttling_rate_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AwsApigatewayv2StageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Stage), &result)
	return &result
}
