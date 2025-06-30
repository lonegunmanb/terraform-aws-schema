package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcEndpointAssociations = `{
  "block": {
    "attributes": {
      "associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associated_resource_accessibility": "string",
              "associated_resource_arn": "string",
              "dns_entry": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "hosted_zone_id": "string"
                  }
                ]
              ],
              "id": "string",
              "private_dns_entry": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "hosted_zone_id": "string"
                  }
                ]
              ],
              "resource_configuration_group_arn": "string",
              "service_network_arn": "string",
              "service_network_name": "string",
              "tags": [
                "map",
                "string"
              ]
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

func AwsVpcEndpointAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcEndpointAssociations), &result)
	return &result
}
