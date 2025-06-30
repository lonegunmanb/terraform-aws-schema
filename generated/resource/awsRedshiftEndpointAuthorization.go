package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftEndpointAuthorization = `{
  "block": {
    "attributes": {
      "account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allowed_all_vpcs": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "grantee": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "grantor": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftEndpointAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftEndpointAuthorization), &result)
	return &result
}
