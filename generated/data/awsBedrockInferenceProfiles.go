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
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
