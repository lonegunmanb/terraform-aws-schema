package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodebuildFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "base_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "compute_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fleet_service_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "overflow_behavior": {
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "context": "string",
              "message": "string",
              "status_code": "string"
            }
          ]
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "compute_configuration": {
        "block": {
          "attributes": {
            "disk": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machine_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "vcpu": {
              "computed": true,
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
      "scaling_configuration": {
        "block": {
          "attributes": {
            "desired_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scaling_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "target_tracking_scaling_configs": {
              "block": {
                "attributes": {
                  "metric_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
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
            },
            "vpc_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodebuildFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodebuildFleet), &result)
	return &result
}
