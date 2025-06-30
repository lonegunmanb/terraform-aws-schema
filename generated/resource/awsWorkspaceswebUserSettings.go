package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspaceswebUserSettings = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "associated_portal_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "copy_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "customer_managed_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deep_link_allowed": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disconnect_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "download_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idle_disconnect_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "paste_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "print_allowed": {
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
      "upload_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cookie_synchronization_configuration": {
        "block": {
          "block_types": {
            "allowlist": {
              "block": {
                "attributes": {
                  "domain": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "blocklist": {
              "block": {
                "attributes": {
                  "domain": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
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
      },
      "toolbar_configuration": {
        "block": {
          "attributes": {
            "hidden_toolbar_items": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_display_resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "toolbar_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "visual_mode": {
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
  "version": 0
}`

func AwsWorkspaceswebUserSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspaceswebUserSettings), &result)
	return &result
}
