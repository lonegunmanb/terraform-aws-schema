package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamOutboundWebIdentityFederation = `{
  "block": {
    "attributes": {
      "issuer_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamOutboundWebIdentityFederationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamOutboundWebIdentityFederation), &result)
	return &result
}
