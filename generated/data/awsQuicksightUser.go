package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightUser = `{
  "block": {
    "attributes": {
      "active": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email": {
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
      "identity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_id": {
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
      "user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsQuicksightUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightUser), &result)
	return &result
}
