package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMacie2OrganizationConfiguration = `{
  "block": {
    "attributes": {
      "auto_enable": {
        "description": "Whether to enable Amazon Macie automatically for accounts that are added to the organization in AWS Organizations",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMacie2OrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMacie2OrganizationConfiguration), &result)
	return &result
}
