package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsResiliencehubResiliencyPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_location_constraint": {
        "computed": true,
        "description": "Specifies a high-level geographical location constraint for where resilience policy data can be stored.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "The description for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "estimated_cost_tier": {
        "computed": true,
        "description": "Specifies the estimated cost tier of the resiliency policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the policy.",
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
      },
      "tier": {
        "description": "The tier for the resiliency policy, ranging from the highest severity (MissionCritical) to lowest (NonCritical).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "policy": {
        "block": {
          "block_types": {
            "az": {
              "block": {
                "attributes": {
                  "rpo": {
                    "description": "Recovery Point Objective (RPO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rto": {
                    "description": "Recovery Time Objective (RTO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The RTO and RPO target to measure resiliency for potential availability zone disruptions.",
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "hardware": {
              "block": {
                "attributes": {
                  "rpo": {
                    "description": "Recovery Point Objective (RPO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rto": {
                    "description": "Recovery Time Objective (RTO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The RTO and RPO target to measure resiliency for potential infrastructure disruptions.",
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "region": {
              "block": {
                "attributes": {
                  "rpo": {
                    "description": "Recovery Point Objective (RPO) as a Go duration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rto": {
                    "description": "Recovery Time Objective (RTO) as a Go duration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The RTO and RPO target to measure resiliency for potential region disruptions.",
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "software": {
              "block": {
                "attributes": {
                  "rpo": {
                    "description": "Recovery Point Objective (RPO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rto": {
                    "description": "Recovery Time Objective (RTO) as a Go duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The RTO and RPO target to measure resiliency for potential application disruptions.",
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            }
          },
          "description": "The resiliency failure policy.",
          "description_kind": "plain"
        },
        "nesting_mode": "single"
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

func AwsResiliencehubResiliencyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsResiliencehubResiliencyPolicy), &result)
	return &result
}
