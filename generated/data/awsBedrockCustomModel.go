package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockCustomModel = `{
  "block": {
    "attributes": {
      "base_model_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hyperparameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
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
      "model_kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "model_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "model_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "output_data_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "s3_uri": "string"
            }
          ]
        ]
      },
      "training_data_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "s3_uri": "string"
            }
          ]
        ]
      },
      "training_metrics": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "training_loss": "number"
            }
          ]
        ]
      },
      "validation_data_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "validator": [
                "list",
                [
                  "object",
                  {
                    "s3_uri": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "validation_metrics": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "validation_loss": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockCustomModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockCustomModel), &result)
	return &result
}
