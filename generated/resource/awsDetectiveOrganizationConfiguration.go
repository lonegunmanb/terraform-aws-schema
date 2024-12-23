package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDetectiveOrganizationConfiguration = `{
  "block": {
    "attributes": {
      "auto_enable": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "graph_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDetectiveOrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDetectiveOrganizationConfiguration), &result)
	return &result
}
