package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGrafanaWorkspaceServiceAccountToken = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "seconds_to_live": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "service_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_account_token_id": {
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

func AwsGrafanaWorkspaceServiceAccountTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGrafanaWorkspaceServiceAccountToken), &result)
	return &result
}
