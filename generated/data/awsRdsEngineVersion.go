package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsEngineVersion = `{
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
        "required": true,
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
      "include_all": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "supported_character_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supported_feature_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supported_modes": {
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
      "supports_certificate_rotation_without_restart": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_global_databases": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_integrations": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_limitless_database": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_local_write_forwarding": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_log_exports_to_cloudwatch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_parallel_query": {
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

func AwsRdsEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsEngineVersion), &result)
	return &result
}
