package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNeptuneEngineVersion = `{
  "block": {
    "attributes": {
      "default_character_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "exportable_log_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "has_major_target": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "has_minor_target": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "parameter_group_family": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_major_targets": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_upgrade_targets": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "supported_character_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supported_timezones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supports_global_databases": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_log_exports_to_cloudwatch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_read_replica": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "valid_major_targets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "valid_minor_targets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "valid_upgrade_targets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_actual": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNeptuneEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNeptuneEngineVersion), &result)
	return &result
}
