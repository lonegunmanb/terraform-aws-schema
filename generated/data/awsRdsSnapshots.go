package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsSnapshots = `{
  "block": {
    "attributes": {
      "db_instance_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_public": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_shared": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshots": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocated_storage": "number",
              "availability_zone": "string",
              "db_instance_identifier": "string",
              "db_snapshot_arn": "string",
              "db_snapshot_identifier": "string",
              "encrypted": "bool",
              "engine": "string",
              "engine_version": "string",
              "iops": "number",
              "kms_key_id": "string",
              "license_model": "string",
              "option_group_name": "string",
              "original_snapshot_create_time": "string",
              "port": "number",
              "snapshot_create_time": "string",
              "snapshot_type": "string",
              "source_db_snapshot_identifier": "string",
              "source_region": "string",
              "status": "string",
              "storage_type": "string",
              "tag_list": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsSnapshotsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsSnapshots), &result)
	return &result
}
