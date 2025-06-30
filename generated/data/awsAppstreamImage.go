package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppstreamImage = `{
  "block": {
    "attributes": {
      "applications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_block_arn": "string",
              "arn": "string",
              "created_time": "string",
              "description": "string",
              "display_name": "string",
              "enabled": "bool",
              "icon_s3_location": [
                "list",
                [
                  "object",
                  {
                    "s3_bucket": "string",
                    "s3_key": "string"
                  }
                ]
              ],
              "icon_url": "string",
              "instance_families": [
                "list",
                "string"
              ],
              "launch_parameters": "string",
              "launch_path": "string",
              "metadata": [
                "map",
                "string"
              ],
              "name": "string",
              "platforms": [
                "list",
                "string"
              ],
              "working_directory": "string"
            }
          ]
        ]
      },
      "appstream_agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_builder_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_builder_supported": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "image_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_fleet": "bool",
              "allow_image_builder": "bool"
            }
          ]
        ]
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_base_image_released_date": {
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
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_change_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string",
              "message": "string"
            }
          ]
        ]
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppstreamImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppstreamImage), &result)
	return &result
}
