package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayIntegration = `{
  "block": {
    "attributes": {
      "cache_key_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cache_namespace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "content_handling": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_http_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_target": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "passthrough_behavior": {
        "computed": true,
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
      "resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "response_transfer_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rest_api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timeout_milliseconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "tls_config": {
        "block": {
          "attributes": {
            "insecure_skip_verification": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AwsApiGatewayIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayIntegration), &result)
	return &result
}
