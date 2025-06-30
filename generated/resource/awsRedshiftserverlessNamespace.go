package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftserverlessNamespace = `{
  "block": {
    "attributes": {
      "admin_password_secret_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "admin_password_secret_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "admin_user_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "admin_user_password_wo": {
        "description_kind": "plain",
        "optional": true,
        "type": "string",
        "write_only": true
      },
      "admin_user_password_wo_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "admin_username": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_iam_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_roles": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_exports": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "manage_admin_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "namespace_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "namespace_name": {
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
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftserverlessNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftserverlessNamespace), &result)
	return &result
}
