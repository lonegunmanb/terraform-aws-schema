package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedaccessTrustProvider = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "device_trust_provider_type": {
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
      "policy_reference_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "trust_provider_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_trust_provider_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "device_options": {
        "block": {
          "attributes": {
            "tenant_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "native_application_oidc_options": {
        "block": {
          "attributes": {
            "authorization_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_secret": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "issuer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_signing_key_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scope": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_info_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oidc_options": {
        "block": {
          "attributes": {
            "authorization_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_secret": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "issuer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scope": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_info_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sse_specification": {
        "block": {
          "attributes": {
            "customer_managed_key_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedaccessTrustProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedaccessTrustProvider), &result)
	return &result
}
