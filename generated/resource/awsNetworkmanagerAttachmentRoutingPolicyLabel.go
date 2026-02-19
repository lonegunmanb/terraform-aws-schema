package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmanagerAttachmentRoutingPolicyLabel = `{
  "block": {
    "attributes": {
      "attachment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "core_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_policy_label": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkmanagerAttachmentRoutingPolicyLabelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmanagerAttachmentRoutingPolicyLabel), &result)
	return &result
}
