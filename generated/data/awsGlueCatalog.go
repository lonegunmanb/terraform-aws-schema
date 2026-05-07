package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueCatalog = `{
  "block": {
    "attributes": {
      "allow_full_table_external_data_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_properties": [
                "map",
                "string"
              ],
              "data_lake_access_properties": [
                "list",
                [
                  "object",
                  {
                    "catalog_type": "string",
                    "data_lake_access": "bool",
                    "data_transfer_role": "string",
                    "kms_key": "string",
                    "managed_workgroup_name": "string",
                    "managed_workgroup_status": "string",
                    "redshift_database_name": "string",
                    "status_message": "string"
                  }
                ]
              ],
              "iceberg_optimization_properties": [
                "list",
                [
                  "object",
                  {
                    "compaction": [
                      "map",
                      "string"
                    ],
                    "orphan_file_deletion": [
                      "map",
                      "string"
                    ],
                    "retention": [
                      "map",
                      "string"
                    ],
                    "role_arn": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
                "list",
                "string"
              ],
              "principal": [
                "list",
                [
                  "object",
                  {
                    "data_lake_principal_identifier": "string"
                  }
                ]
              ]
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
                "list",
                "string"
              ],
              "principal": [
                "list",
                [
                  "object",
                  {
                    "data_lake_principal_identifier": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "federated_catalog": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_name": "string",
              "connection_type": "string",
              "identifier": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
        "type": [
          "map",
          "string"
        ]
      },
      "target_redshift_catalog": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "catalog_arn": "string"
            }
          ]
        ]
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueCatalog), &result)
	return &result
}
