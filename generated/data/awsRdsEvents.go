package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsEvents = `{
  "block": {
    "attributes": {
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_categories": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "events": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "date": "string",
              "event_categories": [
                "list",
                "string"
              ],
              "message": "string",
              "source_arn": "string",
              "source_identifier": "string",
              "source_type": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsEventsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsEvents), &result)
	return &result
}
