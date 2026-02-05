package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketAcl = `{
  "block": {
    "attributes": {
      "acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expected_bucket_owner": {
        "deprecated": true,
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
      }
    },
    "block_types": {
      "access_control_policy": {
        "block": {
          "block_types": {
            "grant": {
              "block": {
                "attributes": {
                  "permission": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "grantee": {
                    "block": {
                      "attributes": {
                        "display_name": {
                          "computed": true,
                          "deprecated": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "email_address": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "uri": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "set"
            },
            "owner": {
              "block": {
                "attributes": {
                  "display_name": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
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

func AwsS3BucketAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketAcl), &result)
	return &result
}
