package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConfigDeliveryChannel = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      },
      "s3_bucket_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_key_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "snapshot_delivery_properties": {
        "block": {
          "attributes": {
            "delivery_frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConfigDeliveryChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConfigDeliveryChannel), &result)
	return &result
}
