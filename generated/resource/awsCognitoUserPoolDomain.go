package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPoolDomain = `{
  "block": {
    "attributes": {
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cloudfront_distribution": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudfront_distribution_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudfront_distribution_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
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
      "managed_login_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AwsCognitoUserPoolDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserPoolDomain), &result)
	return &result
}
