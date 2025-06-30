package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRumMetricsDestination = `{
  "block": {
    "attributes": {
      "app_monitor_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRumMetricsDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRumMetricsDestination), &result)
	return &result
}
