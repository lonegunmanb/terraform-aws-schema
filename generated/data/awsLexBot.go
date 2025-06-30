package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLexBot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "checksum": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "child_directed": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detect_sentiment": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_model_improvements": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "failure_reason": {
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
      "idle_session_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nlu_intent_confidence_threshold": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "voice_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLexBotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLexBot), &result)
	return &result
}
