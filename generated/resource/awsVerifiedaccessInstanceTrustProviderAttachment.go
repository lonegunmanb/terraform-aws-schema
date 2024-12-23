package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedaccessInstanceTrustProviderAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "verifiedaccess_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "verifiedaccess_trust_provider_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedaccessInstanceTrustProviderAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedaccessInstanceTrustProviderAttachment), &result)
	return &result
}
