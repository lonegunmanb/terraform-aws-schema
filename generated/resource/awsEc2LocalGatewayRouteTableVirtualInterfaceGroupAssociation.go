package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_gateway_route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_gateway_virtual_interface_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
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
      "state": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2LocalGatewayRouteTableVirtualInterfaceGroupAssociation), &result)
	return &result
}
