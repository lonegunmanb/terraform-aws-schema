package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheServiceUpdateActions = `{
  "block": {
    "attributes": {
      "cache_cluster_id": {
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
      "replication_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_update_status": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "update_actions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cache_cluster_id": "string",
              "engine": "string",
              "estimated_update_time": "string",
              "recommended_apply_by_date": "string",
              "release_date": "string",
              "replication_group_id": "string",
              "service_update_name": "string",
              "service_update_severity": "string",
              "service_update_status": "string",
              "service_update_type": "string",
              "update_action_status": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheServiceUpdateActionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheServiceUpdateActions), &result)
	return &result
}
