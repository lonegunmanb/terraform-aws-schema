package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Integration = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_handling_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_response_selection_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "integration_subtype": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "passthrough_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payload_format_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "request_templates": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "template_selection_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_milliseconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "response_parameters": {
        "block": {
          "attributes": {
            "mappings": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "status_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "tls_config": {
        "block": {
          "attributes": {
            "server_name_to_verify": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2IntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Integration), &result)
	return &result
}
