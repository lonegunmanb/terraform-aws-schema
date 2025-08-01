package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3AccessPoint = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_origin": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_access_block_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "block_public_acls": "bool",
              "block_public_policy": "bool",
              "ignore_public_acls": "bool",
              "restrict_public_buckets": "bool"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3AccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3AccessPoint), &result)
	return &result
}
