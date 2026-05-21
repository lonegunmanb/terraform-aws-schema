package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubAutomationRuleV2 = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the automation rule.",
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
      "rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description": "The name of the automation rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_order": {
        "description": "The priority of the rule (lower values = higher priority).",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_status": {
        "computed": true,
        "description": "The status of the rule: ENABLED or DISABLED.",
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
      }
    },
    "block_types": {
      "action": {
        "block": {
          "attributes": {
            "type": {
              "description": "The action type: FINDING_FIELDS_UPDATE or EXTERNAL_INTEGRATION.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "external_integration_configuration": {
              "block": {
                "attributes": {
                  "connector_arn": {
                    "description": "The ARN of the connector.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Settings for external integration actions.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "finding_fields_update": {
              "block": {
                "attributes": {
                  "comment": {
                    "description": "A comment for the finding.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "severity_id": {
                    "description": "The severity ID to assign.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status_id": {
                    "description": "The status ID to assign.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Settings for updating finding fields.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Actions to take when the rule matches. Maximum of 1 action.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "criteria": {
        "block": {
          "attributes": {
            "ocsf_finding_criteria_json": {
              "description": "JSON-encoded OCSF finding criteria for the rule.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Filtering type and configuration of the automation rule.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubAutomationRuleV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubAutomationRuleV2), &result)
	return &result
}
