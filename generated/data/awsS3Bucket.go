package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3Bucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bucket_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_regional_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
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
      "website_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "website_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3Bucket), &result)
	return &result
}
