package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointsmsvoicev2Keyword = `{
  "block": {
    "attributes": {
      "keyword": {
        "description": "Keyword to configure. 1-30 characters, upper-case, and cannot start or end with a space.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keyword_action": {
        "computed": true,
        "description": "Action to perform when the keyword is received.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keyword_message": {
        "description": "Message to send when the keyword is received.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origination_identity_arn": {
        "description": "ARN of the origination identity (phone number or pool) to attach the keyword to.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsPinpointsmsvoicev2KeywordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointsmsvoicev2Keyword), &result)
	return &result
}
