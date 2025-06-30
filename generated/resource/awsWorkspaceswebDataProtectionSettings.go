package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspaceswebDataProtectionSettings = `{
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
      "customer_managed_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_protection_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
      }
    },
    "block_types": {
      "inline_redaction_configuration": {
        "block": {
          "attributes": {
            "global_confidence_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "global_enforced_urls": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "global_exempt_urls": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "inline_redaction_pattern": {
              "block": {
                "attributes": {
                  "built_in_pattern_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "confidence_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enforced_urls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exempt_urls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "custom_pattern": {
                    "block": {
                      "attributes": {
                        "keyword_regex": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "pattern_description": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "pattern_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "pattern_regex": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "redaction_place_holder": {
                    "block": {
                      "attributes": {
                        "redaction_place_holder_text": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redaction_place_holder_type": {
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

func AwsWorkspaceswebDataProtectionSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspaceswebDataProtectionSettings), &result)
	return &result
}
