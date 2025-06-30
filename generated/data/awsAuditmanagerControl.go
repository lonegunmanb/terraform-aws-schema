package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAuditmanagerControl = `{
  "block": {
    "attributes": {
      "action_plan_instructions": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "action_plan_title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "control_mapping_sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "source_description": "string",
              "source_frequency": "string",
              "source_id": "string",
              "source_keyword": [
                "list",
                [
                  "object",
                  {
                    "keyword_input_type": "string",
                    "keyword_value": "string"
                  }
                ]
              ],
              "source_name": "string",
              "source_set_up_option": "string",
              "source_type": "string",
              "troubleshooting_text": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
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
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "testing_information": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAuditmanagerControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAuditmanagerControl), &result)
	return &result
}
