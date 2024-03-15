package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMedialiveInput = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attached_channels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "destinations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip": "string",
              "port": "string",
              "url": "string",
              "vpc": [
                "list",
                [
                  "object",
                  {
                    "availability_zone": "string",
                    "network_interface_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "input_devices": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string"
            }
          ]
        ]
      },
      "input_partner_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "input_source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "media_connect_flows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "flow_arn": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "password_param": "string",
              "url": "string",
              "username": "string"
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
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMedialiveInputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMedialiveInput), &result)
	return &result
}
