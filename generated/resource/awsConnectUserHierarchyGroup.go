package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectUserHierarchyGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hierarchy_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hierarchy_path": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "level_five": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_four": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_one": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_three": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_two": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "level_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_group_id": {
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectUserHierarchyGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectUserHierarchyGroup), &result)
	return &result
}
