package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsCertificate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_override": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "customer_override_valid_till": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_for_new_launches": {
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
      "latest_valid_till": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_from": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_till": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsCertificate), &result)
	return &result
}
