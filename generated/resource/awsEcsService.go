package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsService = `{
  "block": {
    "attributes": {
      "cluster": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_maximum_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "deployment_minimum_healthy_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "desired_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_ecs_managed_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_new_deployment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_check_grace_period_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "iam_role": {
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
      "launch_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "propagate_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduling_strategy": {
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
      "task_definition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "capacity_provider_strategy": {
        "block": {
          "attributes": {
            "base": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "capacity_provider": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "deployment_controller": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "load_balancer": {
        "block": {
          "attributes": {
            "container_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "elb_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_group_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "network_configuration": {
        "block": {
          "attributes": {
            "assign_public_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnets": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ordered_placement_strategy": {
        "block": {
          "attributes": {
            "field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      },
      "placement_constraints": {
        "block": {
          "attributes": {
            "expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "set"
      },
      "placement_strategy": {
        "block": {
          "attributes": {
            "field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "set"
      },
      "service_registries": {
        "block": {
          "attributes": {
            "container_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "registry_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
            "delete": {
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

func AwsEcsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsService), &result)
	return &result
}