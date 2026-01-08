package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftIdcApplication = `{
  "block": {
    "attributes": {
      "application_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_instance_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_managed_application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_namespace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redshift_idc_application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "redshift_idc_application_name": {
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
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "authorized_token_issuer": {
        "block": {
          "attributes": {
            "authorized_audiences_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "trusted_token_issuer_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "service_integration": {
        "block": {
          "block_types": {
            "lake_formation": {
              "block": {
                "block_types": {
                  "lake_formation_query": {
                    "block": {
                      "attributes": {
                        "authorization": {
                          "description_kind": "plain",
                          "required": true,
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
            },
            "redshift": {
              "block": {
                "block_types": {
                  "connect": {
                    "block": {
                      "attributes": {
                        "authorization": {
                          "description_kind": "plain",
                          "required": true,
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
            },
            "s3_access_grants": {
              "block": {
                "block_types": {
                  "read_write_access": {
                    "block": {
                      "attributes": {
                        "authorization": {
                          "description_kind": "plain",
                          "required": true,
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
  "version": 0
}`

func AwsRedshiftIdcApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftIdcApplication), &result)
	return &result
}
