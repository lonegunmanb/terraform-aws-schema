package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsArczonalshiftZonalAutoshiftConfiguration = `{
  "block": {
    "attributes": {
      "allowed_windows": {
        "description": "List of time windows during which practice runs are allowed, in the format ` + "`" + `Day:HH:MM-Day:HH:MM` + "`" + ` (e.g., ` + "`" + `Mon:09:00-Mon:17:00` + "`" + `). Cannot be used together with ` + "`" + `blocked_windows` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "blocked_dates": {
        "description": "List of dates when practice runs should not be started, in the format ` + "`" + `YYYY-MM-DD` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "blocked_windows": {
        "description": "List of time windows during which practice runs should not be started, in the format ` + "`" + `Day:HH:MM-Day:HH:MM` + "`" + ` (e.g., ` + "`" + `Mon:00:00-Mon:08:00` + "`" + `). Cannot be used together with ` + "`" + `allowed_windows` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
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
      "resource_arn": {
        "description": "The ARN of the managed resource to configure zonal autoshift for (e.g., an Application Load Balancer). Changing this creates a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zonal_autoshift_status": {
        "description": "The status of zonal autoshift. Valid values: ` + "`" + `ENABLED` + "`" + `, ` + "`" + `DISABLED` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "blocking_alarms": {
        "block": {
          "attributes": {
            "alarm_identifier": {
              "description": "ARN of the CloudWatch alarm.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "Type of control condition. Valid value: ` + "`" + `CLOUDWATCH` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of CloudWatch alarms that can block practice runs when in alarm state.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "outcome_alarms": {
        "block": {
          "attributes": {
            "alarm_identifier": {
              "description": "ARN of the CloudWatch alarm.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "Type of control condition. Valid value: ` + "`" + `CLOUDWATCH` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of CloudWatch alarms monitored during practice runs. These alarms help determine the health of your application during zonal shifts.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsArczonalshiftZonalAutoshiftConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsArczonalshiftZonalAutoshiftConfiguration), &result)
	return &result
}
