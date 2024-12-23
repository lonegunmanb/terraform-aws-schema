package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayDeployment = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description_kind": "plain",
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
      "rest_api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_description": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_name": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "triggers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "canary_settings": {
        "block": {
          "attributes": {
            "percent_traffic": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stage_variable_overrides": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "use_stage_cache": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "deprecated": true,
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

func AwsApiGatewayDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayDeployment), &result)
	return &result
}
