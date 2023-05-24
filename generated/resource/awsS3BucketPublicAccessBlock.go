package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketPublicAccessBlock = `{
  "block": {
    "attributes": {
      "block_public_acls": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "block_public_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bucket": {
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
      "ignore_public_acls": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "restrict_public_buckets": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketPublicAccessBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketPublicAccessBlock), &result)
	return &result
}
