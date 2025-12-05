package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFlowLog = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deliver_cross_account_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eni_id": {
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
      "log_destination": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_destination_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_format": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_aggregation_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regional_nat_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_id": {
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "traffic_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "destination_options": {
        "block": {
          "attributes": {
            "file_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hive_compatible_partitions": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "per_hour_partition": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
  "version": 1
}`

func AwsFlowLogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFlowLog), &result)
	return &result
}
