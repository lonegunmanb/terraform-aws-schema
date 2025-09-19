package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbNetworkPeeringConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Created time of the odb network peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Display name of the odb network peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Network Peering Connection identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "ARN of the odb network peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_peering_connection_type": {
        "computed": true,
        "description": "Type of the odb peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_network_arn": {
        "computed": true,
        "description": "ARN of the peer network peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description": "Progress of the odb network peering connection.",
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
      "status": {
        "computed": true,
        "description": "Status of the odb network peering connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Status of the odb network peering connection.",
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

func AwsOdbNetworkPeeringConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbNetworkPeeringConnection), &result)
	return &result
}
