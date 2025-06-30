package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2NetworkInsightsPath = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "filter_at_destination": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "destination_address": "string",
              "destination_port_range": [
                "list",
                [
                  "object",
                  {
                    "from_port": "number",
                    "to_port": "number"
                  }
                ]
              ],
              "source_address": "string",
              "source_port_range": [
                "list",
                [
                  "object",
                  {
                    "from_port": "number",
                    "to_port": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "filter_at_source": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "destination_address": "string",
              "destination_port_range": [
                "list",
                [
                  "object",
                  {
                    "from_port": "number",
                    "to_port": "number"
                  }
                ]
              ],
              "source_address": "string",
              "source_port_range": [
                "list",
                [
                  "object",
                  {
                    "from_port": "number",
                    "to_port": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_insights_path_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_ip": {
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
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2NetworkInsightsPathSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2NetworkInsightsPath), &result)
	return &result
}
