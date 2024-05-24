package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcEndpointPrivateDns = `{
  "block": {
    "attributes": {
      "private_dns_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "vpc_endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcEndpointPrivateDnsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcEndpointPrivateDns), &result)
	return &result
}
