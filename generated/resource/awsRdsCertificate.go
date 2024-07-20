package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsCertificate = `{
  "block": {
    "attributes": {
      "certificate_identifier": {
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

func AwsRdsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsCertificate), &result)
	return &result
}
