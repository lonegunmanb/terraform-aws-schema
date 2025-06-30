package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessVpcEndpoint = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description": "The date the endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the endpoint.",
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
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of the security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description": "The IDs of the subnets from which you access OpenSearch Serverless.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_endpoint_id": {
        "description": "The unique identifier of the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC from which you access OpenSearch Serverless.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessVpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessVpcEndpoint), &result)
	return &result
}
