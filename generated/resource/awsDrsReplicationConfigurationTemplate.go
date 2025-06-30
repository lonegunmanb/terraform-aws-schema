package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDrsReplicationConfigurationTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associate_default_security_group": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "auto_replicate_new_disks": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth_throttling": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "create_public_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "data_plane_routing": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_large_staging_disk_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ebs_encryption": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ebs_encryption_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
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
      "replication_server_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_servers_security_groups_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "staging_area_subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "staging_area_tags": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "use_dedicated_replication_server": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      }
    },
    "block_types": {
      "pit_policy": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "interval": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "retention_duration": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "rule_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "units": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDrsReplicationConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDrsReplicationConfigurationTemplate), &result)
	return &result
}
