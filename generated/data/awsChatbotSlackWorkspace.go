package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsChatbotSlackWorkspace = `{
  "block": {
    "attributes": {
      "slack_team_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "slack_team_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsChatbotSlackWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsChatbotSlackWorkspace), &result)
	return &result
}
