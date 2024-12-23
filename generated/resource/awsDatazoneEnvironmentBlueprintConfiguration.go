package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatazoneEnvironmentBlueprintConfiguration = `{
  "block": {
    "attributes": {
      "domain_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled_regions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "environment_blueprint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manage_access_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioning_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regional_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          [
            "map",
            "string"
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDatazoneEnvironmentBlueprintConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatazoneEnvironmentBlueprintConfiguration), &result)
	return &result
}
