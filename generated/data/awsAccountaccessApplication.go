package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountaccessApplication = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_instance_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_source": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_center": [
                "list",
                [
                  "object",
                  {
                    "application_arn": "string",
                    "instance_arn": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tenant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAccountaccessApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountaccessApplication), &result)
	return &result
}
