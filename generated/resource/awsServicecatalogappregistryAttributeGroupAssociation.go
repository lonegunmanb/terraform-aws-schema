package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogappregistryAttributeGroupAssociation = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "ID of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "attribute_group_id": {
        "description": "ID of the attribute group to associate with the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AwsServicecatalogappregistryAttributeGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogappregistryAttributeGroupAssociation), &result)
	return &result
}
