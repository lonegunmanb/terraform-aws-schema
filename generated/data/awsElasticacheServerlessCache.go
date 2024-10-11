package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheServerlessCache = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cache_usage_limits": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "data_storage": [
              "object",
              {
                "maximum": "number",
                "minimum": "number",
                "unit": "string"
              }
            ],
            "ecpu_per_second": [
              "object",
              {
                "maximum": "number",
                "minimum": "number"
              }
            ]
          }
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "daily_snapshot_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "address": "string",
            "port": "number"
          }
        ]
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "full_engine_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "major_engine_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reader_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "object",
          {
            "address": "string",
            "port": "number"
          }
        ]
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheServerlessCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheServerlessCache), &result)
	return &result
}
