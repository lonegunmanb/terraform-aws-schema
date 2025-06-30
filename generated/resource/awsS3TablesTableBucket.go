package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3TablesTableBucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "kms_key_arn": "string",
            "sse_algorithm": "string"
          }
        ]
      },
      "maintenance_configuration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "object",
          {
            "iceberg_unreferenced_file_removal": [
              "object",
              {
                "settings": [
                  "object",
                  {
                    "non_current_days": "number",
                    "unreferenced_days": "number"
                  }
                ],
                "status": "string"
              }
            ]
          }
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3TablesTableBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3TablesTableBucket), &result)
	return &result
}
