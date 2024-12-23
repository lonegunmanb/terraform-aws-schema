package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockInferenceProfile = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inference_profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inference_profile_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inference_profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "models": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "model_arn": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockInferenceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockInferenceProfile), &result)
	return &result
}
