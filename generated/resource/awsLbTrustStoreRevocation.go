package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbTrustStoreRevocation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revocation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "revocations_s3_bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revocations_s3_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revocations_s3_object_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trust_store_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLbTrustStoreRevocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbTrustStoreRevocation), &result)
	return &result
}