package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamServiceSpecificCredential = `{
  "block": {
    "attributes": {
      "create_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credential_age_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "expiration_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_credential_alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_credential_secret": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_password": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "service_specific_credential_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamServiceSpecificCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamServiceSpecificCredential), &result)
	return &result
}
