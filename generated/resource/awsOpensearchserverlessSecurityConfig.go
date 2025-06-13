package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessSecurityConfig = `{
  "block": {
    "attributes": {
      "config_version": {
        "computed": true,
        "description": "Version of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the security configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "Type of configuration. Must be ` + "`" + `saml` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "saml_options": {
        "block": {
          "attributes": {
            "group_attribute": {
              "description": "Group attribute for this SAML integration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "description": "The XML IdP metadata file generated from your identity provider.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "session_timeout": {
              "computed": true,
              "description": "Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "user_attribute": {
              "description": "User attribute for this SAML integration.",
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

func AwsOpensearchserverlessSecurityConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessSecurityConfig), &result)
	return &result
}
