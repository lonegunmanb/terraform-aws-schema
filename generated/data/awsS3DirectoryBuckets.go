package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3DirectoryBuckets = `{
  "block": {
    "attributes": {
      "arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "buckets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3DirectoryBucketsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3DirectoryBuckets), &result)
	return &result
}
