package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudformationStackInstances = `{
  "block": {
    "attributes": {
      "accounts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "call_as": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_overrides": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
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
      "regions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "retain_stacks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_instance_summaries": {
        "computed": true,
        "description": "List of stack instances created from an organizational unit deployment target. This will only be populated when ` + "`" + `deployment_targets` + "`" + ` is set.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "detailed_status": "string",
              "drift_status": "string",
              "organizational_unit_id": "string",
              "region": "string",
              "stack_id": "string",
              "stack_set_id": "string",
              "status": "string",
              "status_reason": "string"
            }
          ]
        ]
      },
      "stack_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "deployment_targets": {
        "block": {
          "attributes": {
            "account_filter_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "accounts": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "accounts_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organizational_unit_ids": {
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
        "nesting_mode": "list"
      },
      "operation_preferences": {
        "block": {
          "attributes": {
            "concurrency_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure_tolerance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_tolerance_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "region_concurrency_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_order": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AwsCloudformationStackInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudformationStackInstances), &result)
	return &result
}
