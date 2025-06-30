package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessSecurityConfig = `{
  "block": {
    "attributes": {
      "config_version": {
        "computed": true,
        "description": "The version of the security configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
        "computed": true,
        "description": "The date the configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the security configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "The unique identifier of the security configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "The date the configuration was last modified.",
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
      "type": {
        "computed": true,
        "description": "The type of security configuration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "saml_options": {
        "block": {
          "attributes": {
            "group_attribute": {
              "computed": true,
              "description": "Group attribute for this SAML integration.",
              "description_kind": "plain",
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "The XML IdP metadata file generated from your identity provider.",
              "description_kind": "plain",
              "type": "string"
            },
            "session_timeout": {
              "computed": true,
              "description": "Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.",
              "description_kind": "plain",
              "type": "number"
            },
            "user_attribute": {
              "computed": true,
              "description": "User attribute for this SAML integration.",
              "description_kind": "plain",
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

func AwsOpensearchserverlessSecurityConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessSecurityConfig), &result)
	return &result
}
