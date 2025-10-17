package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightAccountSubscription = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_subscription_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "active_directory_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "admin_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "admin_pro_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "authentication_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "author_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "author_pro_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "contact_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "first_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_identity_center_instance_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_email": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reader_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "reader_pro_group": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "realm": {
        "description_kind": "plain",
        "optional": true,
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
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsQuicksightAccountSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightAccountSubscription), &result)
	return &result
}
