package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmPatchBaseline = `{
  "block": {
    "attributes": {
      "approval_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "approve_after_days": "number",
              "approve_until_date": "string",
              "compliance_level": "string",
              "enable_non_security": "bool",
              "patch_filter": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "approved_patches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "approved_patches_compliance_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "approved_patches_enable_non_security": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "available_security_updates_compliance_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_baseline": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_filter": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key": "string",
              "values": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "json": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operating_system": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner": {
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
      "rejected_patches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "rejected_patches_action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configuration": "string",
              "name": "string",
              "products": [
                "list",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmPatchBaselineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmPatchBaseline), &result)
	return &result
}
