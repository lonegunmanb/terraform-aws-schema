package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsCustomKeyStore = `{
  "block": {
    "attributes": {
      "cloud_hsm_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_key_store_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_key_store_type": {
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
      "key_store_password": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xks_proxy_connectivity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xks_proxy_uri_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xks_proxy_uri_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xks_proxy_vpc_endpoint_service_name": {
        "description_kind": "plain",
        "optional": true,
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
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "xks_proxy_authentication_credential": {
        "block": {
          "attributes": {
            "access_key_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "raw_secret_access_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
