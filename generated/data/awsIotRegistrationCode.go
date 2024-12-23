package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotRegistrationCode = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registration_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIotRegistrationCodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotRegistrationCode), &result)
	return &result
}
