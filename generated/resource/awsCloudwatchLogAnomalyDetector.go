package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchLogAnomalyDetector = `{
  "block": {
    "attributes": {
      "anomaly_visibility_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detector_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "evaluation_frequency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_pattern": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_group_arn_list": {
        "description_kind": "plain",
        "required": true,
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchLogAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchLogAnomalyDetector), &result)
	return &result
}
