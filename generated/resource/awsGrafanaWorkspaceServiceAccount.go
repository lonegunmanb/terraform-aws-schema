package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGrafanaWorkspaceServiceAccount = `{
  "block": {
    "attributes": {
      "grafana_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGrafanaWorkspaceServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGrafanaWorkspaceServiceAccount), &result)
	return &result
}
