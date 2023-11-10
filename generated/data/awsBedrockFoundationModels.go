package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockFoundationModels = `{
  "block": {
    "attributes": {
      "by_customization_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "by_inference_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "by_output_modality": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "by_provider": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "model_summaries": {
        "block": {
          "attributes": {
            "customizations_supported": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "inference_types_supported": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "input_modalities": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "model_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "model_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "model_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "output_modalities": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "provider_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "response_streaming_supported": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockFoundationModelsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockFoundationModels), &result)
	return &result
}
