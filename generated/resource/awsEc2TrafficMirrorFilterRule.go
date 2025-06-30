package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TrafficMirrorFilterRule = `{
  "block": {
    "attributes": {
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
      "destination_cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
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
      "rule_action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_number": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "source_cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_direction": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_mirror_filter_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "destination_port_range": {
        "block": {
          "attributes": {
            "from_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_port_range": {
        "block": {
          "attributes": {
            "from_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AwsEc2TrafficMirrorFilterRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TrafficMirrorFilterRule), &result)
	return &result
}
