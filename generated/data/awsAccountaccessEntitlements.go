package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountaccessEntitlements = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entitlements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "entitlement": [
                "list",
                [
                  "object",
                  {
                    "principal_role": [
                      "list",
                      [
                        "object",
                        {
                          "account_id": "string",
                          "account_name": "string",
                          "principal": [
                            "list",
                            [
                              "object",
                              {
                                "identity_center": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "group_id": "string",
                                      "user_id": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "role_arn": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "entitlement_id": "string"
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
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "block_types": {
            "principal_role": {
              "block": {
                "attributes": {
                  "account_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "optional": true,
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

func AwsAccountaccessEntitlementsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountaccessEntitlements), &result)
	return &result
}
