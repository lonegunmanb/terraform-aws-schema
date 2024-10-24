package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderContainerRecipe = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dockerfile_template_data": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dockerfile_template_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent_image": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_override": {
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
      "version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "working_directory": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "component": {
        "block": {
          "attributes": {
            "component_arn": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "instance_configuration": {
        "block": {
          "attributes": {
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "block_device_mapping": {
              "block": {
                "attributes": {
                  "device_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "no_device": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "virtual_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ebs": {
                    "block": {
                      "attributes": {
                        "delete_on_termination": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "encrypted": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iops": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kms_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "snapshot_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "throughput": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_size": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_type": {
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
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target_repository": {
        "block": {
          "attributes": {
            "repository_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsImagebuilderContainerRecipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderContainerRecipe), &result)
	return &result
}
