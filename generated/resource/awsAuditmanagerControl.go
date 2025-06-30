package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAuditmanagerControl = `{
  "block": {
    "attributes": {
      "action_plan_instructions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "action_plan_title": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "testing_information": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "control_mapping_sources": {
        "block": {
          "attributes": {
            "source_description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_keyword": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "keyword_input_type": "string",
                    "keyword_value": "string"
                  }
                ]
              ]
            },
            "source_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_set_up_option": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "troubleshooting_text": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
