package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaLayerVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "code_sha256": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compatible_architectures": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "compatible_runtimes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filename": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "layer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "layer_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "license_info": {
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
      "s3_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_object_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "signing_job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_profile_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "skip_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_code_hash": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_code_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaLayerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaLayerVersion), &result)
	return &result
}
