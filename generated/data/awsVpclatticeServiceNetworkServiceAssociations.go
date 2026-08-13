package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpclatticeServiceNetworkServiceAssociations = `{
  "block": {
    "attributes": {
      "items": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "created_at": "string",
              "created_by": "string",
              "custom_domain_name": "string",
              "dns_entry": [
                "list",
                [
                  "object",
                  {
                    "domain_name": "string",
                    "hosted_zone_id": "string"
                  }
                ]
              ],
              "id": "string",
              "service_arn": "string",
              "service_id": "string",
              "service_name": "string",
              "service_network_arn": "string",
              "service_network_id": "string",
              "service_network_name": "string",
              "status": "string"
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
      "service_identifier": {
        "description": "ID or ARN of the Service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_network_identifier": {
        "description": "ID or ARN of the Service Network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpclatticeServiceNetworkServiceAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpclatticeServiceNetworkServiceAssociations), &result)
	return &result
}
