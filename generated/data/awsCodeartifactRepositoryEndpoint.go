package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodeartifactRepositoryEndpoint = `{
  "block": {
    "attributes": {
      "domain": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_owner": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "format": {
        "description_kind": "plain",
        "required": true,
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
      "repository": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repository_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodeartifactRepositoryEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodeartifactRepositoryEndpoint), &result)
	return &result
}
