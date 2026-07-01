package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointsmsvoicev2Pool = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "Whether deletion protection is enabled. When ` + "`" + `true` + "`" + `, the pool cannot be deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iso_country_code": {
        "description": "Two-character code, in ISO 3166-1 alpha-2 format, for the country or region of the pool. This field is optional for origination identity types that are not country-specific.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "message_type": {
        "description": "Type of message.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "opt_out_list_name": {
        "computed": true,
        "description": "Name of the opt-out list to associate with the pool. Inherited from the initial origination identity when omitted.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origination_identities": {
        "description": "Set of origination identity ARNs to associate with the pool. At least one origination identity is required at creation.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
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
      "self_managed_opt_outs_enabled": {
        "computed": true,
        "description": "Whether the pool relies on self-managed opt-out handling. When ` + "`" + `false` + "`" + `, AWS auto-replies to HELP/STOP requests and manages the opt-out list. Inherited from the initial origination identity when omitted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "shared_routes_enabled": {
        "computed": true,
        "description": "Whether shared routes are enabled for the pool. When ` + "`" + `true` + "`" + `, messages may use shared phone numbers or sender IDs in countries that allow it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "two_way_channel_arn": {
        "description": "ARN of the two-way channel that receives inbound messages.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_channel_role": {
        "description": "ARN of the IAM role that End User Messaging SMS assumes to publish inbound messages to the two-way channel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "two_way_enabled": {
        "computed": true,
        "description": "Whether inbound message reception is enabled for the pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsPinpointsmsvoicev2PoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointsmsvoicev2Pool), &result)
	return &result
}
