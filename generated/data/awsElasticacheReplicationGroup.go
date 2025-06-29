package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheReplicationGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_token_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "automatic_failover_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cluster_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "log_delivery_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "destination": "string",
              "destination_type": "string",
              "log_format": "string",
              "log_type": "string"
            }
          ]
        ]
      },
      "member_clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "multi_az_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "num_cache_clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "num_node_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "primary_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reader_endpoint_address": {
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
      "replicas_per_node_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "replication_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheReplicationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheReplicationGroup), &result)
	return &result
}
