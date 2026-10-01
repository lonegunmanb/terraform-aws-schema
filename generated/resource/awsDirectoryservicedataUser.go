package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDirectoryservicedataUser = `{
  "block": {
    "attributes": {
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "distinguished_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "given_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "realm": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sam_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "surname": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_principal_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDirectoryservicedataUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDirectoryservicedataUser), &result)
	return &result
}
