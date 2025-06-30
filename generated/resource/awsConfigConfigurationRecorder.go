package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConfigConfigurationRecorder = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "recording_group": {
        "block": {
          "attributes": {
            "all_supported": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_global_resource_types": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "resource_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "exclusion_by_resource_types": {
              "block": {
                "attributes": {
                  "resource_types": {
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
            },
            "recording_strategy": {
              "block": {
                "attributes": {
                  "use_only": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "recording_mode": {
        "block": {
          "attributes": {
            "recording_frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "recording_mode_override": {
              "block": {
                "attributes": {
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "recording_frequency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_types": {
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

func AwsConfigConfigurationRecorderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConfigConfigurationRecorder), &result)
	return &result
}
