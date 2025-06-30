package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOutpostsAsset = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_id": {
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
      "rack_elevation": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "rack_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOutpostsAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOutpostsAsset), &result)
	return &result
}
