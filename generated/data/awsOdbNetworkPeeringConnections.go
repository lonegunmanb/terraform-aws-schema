package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbNetworkPeeringConnections = `{
  "block": {
    "attributes": {
      "odb_peering_connections": {
        "computed": true,
        "description": "The list of ODB peering connections. A summary of an ODB peering connection.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "display_name": "string",
              "id": "string",
              "odb_network_arn": "string",
              "peer_network_arn": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOdbNetworkPeeringConnectionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbNetworkPeeringConnections), &result)
	return &result
}
