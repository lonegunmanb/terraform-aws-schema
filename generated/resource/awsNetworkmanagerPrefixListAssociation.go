package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmanagerPrefixListAssociation = `{
  "block": {
    "attributes": {
      "core_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_alias": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkmanagerPrefixListAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmanagerPrefixListAssociation), &result)
	return &result
}
