package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointEmailChannel = `{
  "block": {
    "attributes": {
      "application_id": {
        "deprecated": true,
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_set": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "from_address": {
        "deprecated": true,
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "deprecated": true,
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "messages_per_second": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "number"
      },
      "orchestration_sending_role_arn": {
        "deprecated": true,
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
      },
      "role_arn": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsPinpointEmailChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointEmailChannel), &result)
	return &result
}
