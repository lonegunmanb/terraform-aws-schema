package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPool = `{
  "block": {
    "attributes": {
      "account_recovery_setting": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "recovery_mechanism": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "priority": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "admin_create_user_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_admin_create_user_only": "bool",
              "invite_message_template": [
                "list",
                [
                  "object",
                  {
                    "email_message": "string",
                    "email_subject": "string",
                    "sms_message": "string"
                  }
                ]
              ],
              "unused_account_validity_days": "number"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_verified_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "device_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "challenge_required_on_new_device": "bool",
              "device_only_remembered_on_user_prompt": "bool"
            }
          ]
        ]
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configuration_set": "string",
              "email_sending_account": "string",
              "from": "string",
              "reply_to_email_address": "string",
              "source_arn": "string"
            }
          ]
        ]
      },
      "estimated_number_of_users": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lambda_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_auth_challenge": "string",
              "custom_email_sender": [
                "list",
                [
                  "object",
                  {
                    "lambda_arn": "string",
                    "lambda_version": "string"
                  }
                ]
              ],
              "custom_message": "string",
              "custom_sms_sender": [
                "list",
                [
                  "object",
                  {
                    "lambda_arn": "string",
                    "lambda_version": "string"
                  }
                ]
              ],
              "define_auth_challenge": "string",
              "kms_key_id": "string",
              "post_authentication": "string",
              "post_confirmation": "string",
              "pre_authentication": "string",
              "pre_sign_up": "string",
              "pre_token_generation": "string",
              "pre_token_generation_config": [
                "list",
                [
                  "object",
                  {
                    "lambda_arn": "string",
                    "lambda_version": "string"
                  }
                ]
              ],
              "user_migration": "string",
              "verify_auth_challenge_response": "string"
            }
          ]
        ]
      },
      "last_modified_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mfa_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
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
      "schema_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attribute_data_type": "string",
              "developer_only_attribute": "bool",
              "mutable": "bool",
              "name": "string",
              "number_attribute_constraints": [
                "list",
                [
                  "object",
                  {
                    "max_value": "string",
                    "min_value": "string"
                  }
                ]
              ],
              "required": "bool",
              "string_attribute_constraints": [
                "list",
                [
                  "object",
                  {
                    "max_length": "string",
                    "min_length": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "sms_authentication_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sms_configuration_failure": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sms_verification_message": {
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
      "user_pool_add_ons": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "advanced_security_additional_flows": [
                "list",
                [
                  "object",
                  {
                    "custom_auth_mode": "string"
                  }
                ]
              ],
              "advanced_security_mode": "string"
            }
          ]
        ]
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_pool_tags": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "username_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoUserPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserPool), &result)
	return &result
}
