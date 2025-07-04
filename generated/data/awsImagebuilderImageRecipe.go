package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderImageRecipe = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "block_device_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "device_name": "string",
              "ebs": [
                "list",
                [
                  "object",
                  {
                    "delete_on_termination": "string",
                    "encrypted": "string",
                    "iops": "number",
                    "kms_key_id": "string",
                    "snapshot_id": "string",
                    "throughput": "number",
                    "volume_size": "number",
                    "volume_type": "string"
                  }
                ]
              ],
              "no_device": "string",
              "virtual_name": "string"
            }
          ]
        ]
      },
      "component": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "component_arn": "string",
              "parameter": [
                "set",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent_image": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_data_base64": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "working_directory": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsImagebuilderImageRecipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderImageRecipe), &result)
	return &result
}
