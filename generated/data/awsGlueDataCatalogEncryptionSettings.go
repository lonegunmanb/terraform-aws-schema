package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueDataCatalogEncryptionSettings = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_catalog_encryption_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_password_encryption": [
                "list",
                [
                  "object",
                  {
                    "aws_kms_key_id": "string",
                    "return_connection_password_encrypted": "bool"
                  }
                ]
              ],
              "encryption_at_rest": [
                "list",
                [
                  "object",
                  {
                    "catalog_encryption_mode": "string",
                    "catalog_encryption_service_role": "string",
                    "sse_aws_kms_key_id": "string"
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

func AwsGlueDataCatalogEncryptionSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueDataCatalogEncryptionSettings), &result)
	return &result
}
