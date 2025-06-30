package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLakeformationDataLakeSettings = `{
  "block": {
    "attributes": {
      "admins": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "allow_external_data_filtering": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_full_table_external_data_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "authorized_session_tag_value_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "catalog_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_database_default_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "permissions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "create_table_default_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "permissions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "external_data_filtering_allow_list": {
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
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "read_only_admins": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trusted_resource_owners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLakeformationDataLakeSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLakeformationDataLakeSettings), &result)
	return &result
}
