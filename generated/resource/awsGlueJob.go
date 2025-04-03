package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueJob = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connections": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "default_arguments": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "glue_version": {
        "computed": true,
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
      "job_run_queuing_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "maintenance_window": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_retries": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "non_overridable_arguments": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "number_of_workers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_configuration": {
        "description_kind": "plain",
        "optional": true,
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
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "command": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script_location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "execution_property": {
        "block": {
          "attributes": {
            "max_concurrent_runs": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_property": {
        "block": {
          "attributes": {
            "notify_delay_after": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_control_details": {
        "block": {
          "attributes": {
            "auth_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_token": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "branch": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "folder": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_commit_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "owner": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provider": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "repository": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueJob), &result)
	return &result
}
