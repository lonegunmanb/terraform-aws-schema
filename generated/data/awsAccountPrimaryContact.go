package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountPrimaryContact = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_line_1": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "address_line_2": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "address_line_3": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "city": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "company_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "country_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "district_or_county": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "full_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "phone_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "postal_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_or_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "website_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAccountPrimaryContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountPrimaryContact), &result)
	return &result
}
