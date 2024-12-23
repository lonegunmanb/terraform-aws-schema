package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGuarddutyOrganizationConfigurationFeature = `{
  "block": {
    "attributes": {
      "auto_enable": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "detector_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
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
    "block_types": {
      "additional_configuration": {
        "block": {
          "attributes": {
            "auto_enable": {
              "description_kind": "plain",
              "required": true,
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGuarddutyOrganizationConfigurationFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGuarddutyOrganizationConfigurationFeature), &result)
	return &result
}
