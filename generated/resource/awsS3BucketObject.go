package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketObject = `{
  "block": {
    "attributes": {
      "acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "deprecated": true,
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bucket_key_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cache_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_base64": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_disposition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key": {
        "deprecated": true,
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "object_lock_legal_hold_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_lock_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_lock_retain_until_date": {
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
      "server_side_encryption": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_hash": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "website_redirect": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketObjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketObject), &result)
	return &result
}
