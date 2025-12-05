package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAlbTargetGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn_suffix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_termination": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deregistration_delay": {
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
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_multi_value_headers_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "load_balancer_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "load_balancing_algorithm_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_anomaly_mitigation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing_cross_zone_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "preserve_client_ip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_protocol_v2": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slow_start": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "target_control_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "target_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "health_check": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "healthy_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "matcher": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unhealthy_threshold": {
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
      "stickiness": {
        "block": {
          "attributes": {
            "cookie_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cookie_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
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
      "target_failover": {
        "block": {
          "attributes": {
            "on_deregistration": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "on_unhealthy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "target_group_health": {
        "block": {
          "block_types": {
            "dns_failover": {
              "block": {
                "attributes": {
                  "minimum_healthy_targets_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "minimum_healthy_targets_percentage": {
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
            "unhealthy_state_routing": {
              "block": {
                "attributes": {
                  "minimum_healthy_targets_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minimum_healthy_targets_percentage": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target_health_state": {
        "block": {
          "attributes": {
            "enable_unhealthy_connection_termination": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "unhealthy_draining_interval": {
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
  "version": 0
}`

func AwsAlbTargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAlbTargetGroup), &result)
	return &result
}
