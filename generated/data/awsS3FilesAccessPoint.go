package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3FilesAccessPoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_id": {
        "computed": true,
        "description": "File system ID",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Access point ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Access point name",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AWS account ID of the owner",
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
      "status": {
        "computed": true,
        "description": "Access point status",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "posix_user": {
        "block": {
          "attributes": {
            "gid": {
              "computed": true,
              "description": "POSIX group ID",
              "description_kind": "plain",
              "type": "number"
            },
            "secondary_gids": {
              "computed": true,
              "description": "Secondary POSIX group IDs",
              "description_kind": "plain",
              "type": [
                "set",
                "number"
              ]
            },
            "uid": {
              "computed": true,
              "description": "POSIX user ID",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "root_directory": {
        "block": {
          "attributes": {
            "path": {
              "computed": true,
              "description": "Root directory path",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "creation_permissions": {
              "block": {
                "attributes": {
                  "owner_gid": {
                    "computed": true,
                    "description": "Owner group ID",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "owner_uid": {
                    "computed": true,
                    "description": "Owner user ID",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "permissions": {
                    "computed": true,
                    "description": "POSIX permissions",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
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

func AwsS3FilesAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3FilesAccessPoint), &result)
	return &result
}
