package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderImagePipeline = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_recipe_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_last_run": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_next_run": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_updated": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distribution_configuration_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_role": {
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
      "image_recipe_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "infrastructure_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "block_types": {
      "image_scanning_configuration": {
        "block": {
          "attributes": {
            "image_scanning_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ecr_configuration": {
              "block": {
                "attributes": {
                  "container_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "repository_name": {
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
      "image_tests_configuration": {
        "block": {
          "attributes": {
            "image_tests_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "timeout_minutes": {
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
      "schedule": {
        "block": {
          "attributes": {
            "pipeline_execution_start_condition": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timezone": {
              "computed": true,
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
      "workflow": {
        "block": {
          "attributes": {
            "on_failure": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parallel_group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "workflow_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "parameter": {
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
              "nesting_mode": "set"
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

func AwsImagebuilderImagePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderImagePipeline), &result)
	return &result
}
