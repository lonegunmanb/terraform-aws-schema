package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecretsmanagerRandomPassword = `{
  "block": {
    "attributes": {
      "exclude_characters": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_lowercase": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "exclude_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "exclude_punctuation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "exclude_uppercase": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_space": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "password_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "random_password": {
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
      "require_each_included_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecretsmanagerRandomPasswordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecretsmanagerRandomPassword), &result)
	return &result
}
