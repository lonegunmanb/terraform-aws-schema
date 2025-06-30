package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockFoundationModel = `{
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
        "description_kind": "plain",
        "required": true,
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
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
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
  "version": 0
}`

func AwsBedrockFoundationModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockFoundationModel), &result)
	return &result
}
