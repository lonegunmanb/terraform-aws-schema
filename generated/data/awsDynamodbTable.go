package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDynamodbTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attribute": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "type": "string"
            }
          ]
        ]
      },
      "billing_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "global_secondary_index": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "hash_key": "string",
              "name": "string",
              "non_key_attributes": [
                "list",
                "string"
              ],
              "on_demand_throughput": [
                "list",
                [
                  "object",
                  {
                    "max_read_request_units": "number",
                    "max_write_request_units": "number"
                  }
                ]
              ],
              "projection_type": "string",
              "range_key": "string",
              "read_capacity": "number",
              "warm_throughput": [
                "list",
                [
                  "object",
                  {
                    "read_units_per_second": "number",
                    "write_units_per_second": "number"
                  }
                ]
              ],
              "write_capacity": "number"
            }
          ]
        ]
      },
      "hash_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_secondary_index": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "non_key_attributes": [
                "list",
                "string"
              ],
              "projection_type": "string",
              "range_key": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_demand_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_read_request_units": "number",
              "max_write_request_units": "number"
            }
          ]
        ]
      },
      "point_in_time_recovery": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "recovery_period_in_days": "number"
            }
          ]
        ]
      },
      "range_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "read_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "kms_key_arn": "string",
              "region_name": "string"
            }
          ]
        ]
      },
      "stream_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "stream_label": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_view_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "attribute_name": "string",
              "enabled": "bool"
            }
          ]
        ]
      },
      "warm_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "read_units_per_second": "number",
              "write_units_per_second": "number"
            }
          ]
        ]
      },
      "write_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "server_side_encryption": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
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

func AwsDynamodbTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDynamodbTable), &result)
	return &result
}
