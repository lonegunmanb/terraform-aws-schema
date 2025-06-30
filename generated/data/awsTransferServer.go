package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTransferServer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_type": {
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
      "identity_provider_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "invocation_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocols": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "structured_log_destinations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AwsTransferServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTransferServer), &result)
	return &result
}
