package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3FilesFileSystems = `{
  "block": {
    "attributes": {
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "file_systems": {
        "block": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "File system ARN",
              "description_kind": "plain",
              "type": "string"
            },
            "bucket": {
              "computed": true,
              "description": "S3 bucket ARN",
              "description_kind": "plain",
              "type": "string"
            },
            "creation_time": {
              "computed": true,
              "description": "Creation time",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "File system ID",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
              "description": "KMS key ID for encryption",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "File system name",
              "description_kind": "plain",
              "type": "string"
            },
            "owner_id": {
              "computed": true,
              "description": "AWS account ID of the owner",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "IAM role ARN",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "File system status",
              "description_kind": "plain",
              "type": "string"
            },
            "status_message": {
              "computed": true,
              "description": "Status message",
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
  "version": 0
}`

func AwsS3FilesFileSystemsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3FilesFileSystems), &result)
	return &result
}
