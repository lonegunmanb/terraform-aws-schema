package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsComputeoptimizerRecommendationPreferences = `{
  "block": {
    "attributes": {
      "enhanced_infrastructure_metrics": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inferred_workload_types": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "look_back_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "savings_estimation_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "external_metrics_preference": {
        "block": {
          "attributes": {
            "source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "preferred_resource": {
        "block": {
          "attributes": {
            "exclude_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "include_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "scope": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "utilization_preference": {
        "block": {
          "attributes": {
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "metric_parameters": {
              "block": {
                "attributes": {
                  "headroom": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description_kind": "plain",
                    "optional": true,
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
  "version": 0
}`

func AwsComputeoptimizerRecommendationPreferencesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsComputeoptimizerRecommendationPreferences), &result)
	return &result
}
