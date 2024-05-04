package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentAgentKnowledgeBaseAssociation = `{
  "block": {
    "attributes": {
      "agent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "knowledge_base_state": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockagentAgentKnowledgeBaseAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentAgentKnowledgeBaseAssociation), &result)
	return &result
}
