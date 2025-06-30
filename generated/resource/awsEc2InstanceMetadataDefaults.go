package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2InstanceMetadataDefaults = `{
  "block": {
    "attributes": {
      "http_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_put_response_hop_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "http_tokens": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_metadata_tags": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2InstanceMetadataDefaultsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2InstanceMetadataDefaults), &result)
	return &result
}
