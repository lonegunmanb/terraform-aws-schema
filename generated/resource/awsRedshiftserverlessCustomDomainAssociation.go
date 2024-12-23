package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftserverlessCustomDomainAssociation = `{
  "block": {
    "attributes": {
      "custom_domain_certificate_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_domain_certificate_expiry_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workgroup_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftserverlessCustomDomainAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftserverlessCustomDomainAssociation), &result)
	return &result
}
