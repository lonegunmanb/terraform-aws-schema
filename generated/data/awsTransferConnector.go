package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTransferConnector = `{
  "block": {
    "attributes": {
      "access_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "as2_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "basic_auth_secret_id": "string",
              "compression": "string",
              "encryption_algorithm": "string",
              "local_profile_id": "string",
              "mdn_response": "string",
              "mdn_signing_algorithm": "string",
              "message_subject": "string",
              "partner_profile_id": "string",
              "singing_algorithm": "string"
            }
          ]
        ]
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logging_role": {
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
      "security_policy_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_managed_egress_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sftp_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "trusted_host_keys": [
                "list",
                "string"
              ],
              "user_secret_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTransferConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTransferConnector), &result)
	return &result
}
