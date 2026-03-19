package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayMeteringPolicyEntry = `{
  "block": {
    "attributes": {
      "destination_cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_port_range": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_transit_gateway_attachment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_transit_gateway_attachment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metered_account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_rule_number": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
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
      "source_cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_port_range": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_transit_gateway_attachment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_transit_gateway_attachment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_metering_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2TransitGatewayMeteringPolicyEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayMeteringPolicyEntry), &result)
	return &result
}
