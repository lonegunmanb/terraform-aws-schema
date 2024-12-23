package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockInferenceProfiles = `{
  "block": {
    "attributes": {
      "inference_profile_summaries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "description": "string",
              "inference_profile_arn": "string",
              "inference_profile_id": "string",
              "inference_profile_name": "string",
              "models": [
                "list",
                [
                  "object",
                  {
                    "model_arn": "string"
                  }
                ]
              ],
              "status": "string",
              "type": "string",
              "updated_at": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockInferenceProfilesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockInferenceProfiles), &result)
	return &result
}
