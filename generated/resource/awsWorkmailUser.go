package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkmailUser = `{
  "block": {
    "attributes": {
      "city": {
        "description": "City where the user is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "company": {
        "description": "Company associated with the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "country": {
        "description": "Country where the user is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "department": {
        "description": "Department associated with the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled_date": {
        "computed": true,
        "description": "Timestamp when the user was disabled from WorkMail use.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Display name of the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email": {
        "description": "Primary email address used to register the user with WorkMail.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled_date": {
        "computed": true,
        "description": "Timestamp when the user was enabled for WorkMail use.",
        "description_kind": "plain",
        "type": "string"
      },
      "first_name": {
        "description": "First name of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hidden_from_global_address_list": {
        "computed": true,
        "description": "Whether to hide the user from the global address list.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "identity_provider_identity_store_id": {
        "computed": true,
        "description": "Identity store ID from IAM Identity Center associated with the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_user_id": {
        "description": "User ID from IAM Identity Center associated with the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initials": {
        "description": "Initials of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_title": {
        "description": "Job title of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_name": {
        "description": "Last name of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mailbox_deprovisioned_date": {
        "computed": true,
        "description": "Timestamp when the mailbox was removed for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "mailbox_provisioned_date": {
        "computed": true,
        "description": "Timestamp when the mailbox was created for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Username of the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "office": {
        "description": "Office where the user is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "organization_id": {
        "description": "Identifier of the WorkMail organization where the user is managed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description": "Password to set for the user.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current WorkMail state of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "street": {
        "description": "Street address of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "telephone": {
        "description": "Telephone number of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description": "Identifier of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_role": {
        "computed": true,
        "description": "Role assigned to the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zip_code": {
        "description": "ZIP or postal code of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWorkmailUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkmailUser), &result)
	return &result
}
