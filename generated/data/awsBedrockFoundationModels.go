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
      },
      "model_summaries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "customizations_supported": [
                "set",
                "string"
              ],
              "inference_types_supported": [
                "set",
                "string"
              ],
              "input_modalities": [
                "set",
                "string"
              ],
              "model_arn": "string",
              "model_id": "string",
              "model_name": "string",
              "output_modalities": [
                "set",
                "string"
              ],
              "provider_name": "string",
              "response_streaming_supported": "bool"
            }
          ]
        ]
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
