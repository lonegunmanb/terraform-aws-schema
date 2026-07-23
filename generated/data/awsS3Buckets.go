package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3Buckets = `{
  "block": {
    "attributes": {
      "bucket_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "buckets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bucket_arn": "string",
              "bucket_region": "string",
              "creation_date": "string",
              "name": "string"
            }
          ]
        ]
      },
      "max_buckets": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "prefix": {
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

func AwsS3BucketsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3Buckets), &result)
	return &result
}
