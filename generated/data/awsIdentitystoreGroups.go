package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "display_name": "string",
              "external_ids": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "issuer": "string"
                  }
                ]
              ],
              "group_id": "string",
              "identity_store_id": "string"
            }
          ]
        ]
      },
      "identity_store_id": {
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

func AwsIdentitystoreGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreGroups), &result)
	return &result
}
