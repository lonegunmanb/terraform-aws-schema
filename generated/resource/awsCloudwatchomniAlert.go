package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchomniAlert = `{
  "block": {
    "attributes": {
      "alert_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notifications_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "profile_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "space_id": {
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
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "notification_rule": {
        "block": {
          "block_types": {
            "target": {
              "block": {
                "attributes": {
                  "arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metadata": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trigger": {
              "block": {
                "attributes": {
                  "state_values": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "rule": {
        "block": {
          "block_types": {
            "telemetry_rule": {
              "block": {
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "comparator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "critical_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "threshold_field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "threshold_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "warning_threshold": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "evaluation": {
                    "block": {
                      "attributes": {
                        "interval_seconds": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "pending_duration_seconds": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "recovery_duration_seconds": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "no_data": {
                    "block": {
                      "attributes": {
                        "treat_as": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "query": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "language": {
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
              "nesting_mode": "list"
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

func AwsCloudwatchomniAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchomniAlert), &result)
	return &result
}
