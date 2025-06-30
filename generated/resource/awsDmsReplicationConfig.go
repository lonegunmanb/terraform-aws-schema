package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsReplicationConfig = `{
  "block": {
    "attributes": {
      "arn": {
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
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_config_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_settings": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_replication": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supplemental_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_mappings": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_endpoint_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "compute_config": {
        "block": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dns_name_servers": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_capacity_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_capacity_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "multi_az": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "preferred_maintenance_window": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "replication_subnet_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsDmsReplicationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsReplicationConfig), &result)
	return &result
}
