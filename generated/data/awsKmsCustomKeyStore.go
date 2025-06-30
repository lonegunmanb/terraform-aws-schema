package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsCustomKeyStore = `{
  "block": {
    "attributes": {
      "cloud_hsm_cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_key_store_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_key_store_name": {
        "computed": true,
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
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trust_anchor_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKmsCustomKeyStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKmsCustomKeyStore), &result)
	return &result
}
