package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsReplicationTask = `{
  "block": {
    "attributes": {
      "cdc_start_position": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cdc_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migration_type": {
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
      "replication_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_task_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_task_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_task_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_endpoint_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_replication_task": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_mappings": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_endpoint_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDmsReplicationTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsReplicationTask), &result)
	return &result
}
