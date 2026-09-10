package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountaccessEntitlement = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entitlement_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "entitlement": {
        "block": {
          "block_types": {
            "principal_role": {
              "block": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "account_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "principal": {
                    "block": {
                      "block_types": {
                        "identity_center": {
                          "block": {
                            "attributes": {
                              "group_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAccountaccessEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountaccessEntitlement), &result)
	return &result
}
