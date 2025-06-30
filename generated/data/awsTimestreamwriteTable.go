package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTimestreamwriteTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "magnetic_store_write_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_magnetic_store_writes": "bool",
              "magnetic_store_rejected_data_location": [
                "list",
                [
                  "object",
                  {
                    "s3_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "bucket_name": "string",
                          "encryption_option": "string",
                          "kms_key_id": "string",
                          "object_key_prefix": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "magnetic_store_retention_period_in_days": "number",
              "memory_store_retention_period_in_hours": "number"
            }
          ]
        ]
      },
      "schema": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "composite_partition_key": [
                "list",
                [
                  "object",
                  {
                    "enforcement_in_record": "string",
                    "name": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "table_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTimestreamwriteTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTimestreamwriteTable), &result)
	return &result
}
