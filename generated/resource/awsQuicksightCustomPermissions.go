package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightCustomPermissions = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_permissions_name": {
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
      }
    },
    "block_types": {
      "capabilities": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "add_or_run_anomaly_detection_for_analyses": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "analysis": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "automate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "chat_agent": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_dashboard_email_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_data_sources": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_datasets": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_themes": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_threshold_alerts": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_chat_agents": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_shared_folders": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_spice_dataset": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dashboard": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_csv": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_csv_in_scheduled_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_excel": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_excel_in_scheduled_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_pdf": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_pdf_in_scheduled_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flow": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "include_content_in_scheduled_reports_email": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "knowledge_base": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "perform_flow_ui_task": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "print_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_without_approval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rename_shared_folders": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "research": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_analyses": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_dashboards": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_data_sources": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_datasets": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "space": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subscribe_dashboard_email_reports": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_agent_web_search": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_bedrock_models": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "view_account_spice_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsQuicksightCustomPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightCustomPermissions), &result)
	return &result
}
