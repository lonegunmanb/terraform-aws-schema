package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPoolUiCustomization = `{
  "block": {
    "attributes": {
      "client_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "css": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "css_version": {
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
      "image_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_date": {
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
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoUserPoolUiCustomizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserPoolUiCustomization), &result)
	return &result
}
