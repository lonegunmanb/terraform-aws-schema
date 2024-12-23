package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmonitorProbe = `{
  "block": {
    "attributes": {
      "address_family": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "packet_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "probe_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_arn": {
        "description_kind": "plain",
        "required": true,
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
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkmonitorProbeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmonitorProbe), &result)
	return &result
}
