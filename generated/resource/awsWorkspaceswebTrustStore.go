package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspaceswebTrustStore = `{
  "block": {
    "attributes": {
      "associated_portal_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
        "type": [
          "map",
          "string"
        ]
      },
      "trust_store_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "certificate": {
        "block": {
          "attributes": {
            "body": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "issuer": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "not_valid_after": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "not_valid_before": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subject": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "thumbprint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWorkspaceswebTrustStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspaceswebTrustStore), &result)
	return &result
}
