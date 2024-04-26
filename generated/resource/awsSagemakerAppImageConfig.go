package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerAppImageConfig = `{
  "block": {
    "attributes": {
      "app_image_config_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
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
      "code_editor_app_image_config": {
        "block": {
          "block_types": {
            "container_config": {
              "block": {
                "attributes": {
                  "container_arguments": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_entrypoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_environment_variables": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "file_system_config": {
              "block": {
                "attributes": {
                  "default_gid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "default_uid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mount_path": {
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
      "jupyter_lab_image_config": {
        "block": {
          "block_types": {
            "container_config": {
              "block": {
                "attributes": {
                  "container_arguments": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_entrypoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_environment_variables": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "file_system_config": {
              "block": {
                "attributes": {
                  "default_gid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "default_uid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mount_path": {
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
      "kernel_gateway_image_config": {
        "block": {
          "block_types": {
            "file_system_config": {
              "block": {
                "attributes": {
                  "default_gid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "default_uid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mount_path": {
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
            "kernel_spec": {
              "block": {
                "attributes": {
                  "display_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "min_items": 1,
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

func AwsSagemakerAppImageConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerAppImageConfig), &result)
	return &result
}
