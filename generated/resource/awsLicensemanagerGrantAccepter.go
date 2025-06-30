package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLicensemanagerGrantAccepter = `{
  "block": {
    "attributes": {
      "allowed_operations": {
        "computed": true,
        "description": "Allowed operations for the grant.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "grant_arn": {
        "description": "Amazon Resource Name (ARN) of the grant.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_region": {
        "computed": true,
        "description": "Home Region of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_arn": {
        "computed": true,
        "description": "License ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_arn": {
        "computed": true,
        "description": "Parent ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The grantee principal ARN.",
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
      "status": {
        "computed": true,
        "description": "GrantAccepter status.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "GrantAccepter version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLicensemanagerGrantAccepterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLicensemanagerGrantAccepter), &result)
	return &result
}
