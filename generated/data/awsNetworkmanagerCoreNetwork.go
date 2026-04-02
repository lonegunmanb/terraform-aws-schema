package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmanagerCoreNetwork = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "edges": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "asn": "number",
              "edge_location": "string",
              "inside_cidr_blocks": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "global_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_function_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "edge_locations": [
                "list",
                "string"
              ],
              "name": "string",
              "segments": [
                "list",
                [
                  "object",
                  {
                    "send_to": [
                      "list",
                      "string"
                    ],
                    "send_via": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "segments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "edge_locations": [
                "list",
                "string"
              ],
              "name": "string",
              "shared_segments": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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

func AwsNetworkmanagerCoreNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmanagerCoreNetwork), &result)
	return &result
}
