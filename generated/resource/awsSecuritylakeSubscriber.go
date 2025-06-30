package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecuritylakeSubscriber = `{
  "block": {
    "attributes": {
      "access_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
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
      },
      "resource_share_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_share_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_bucket_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscriber_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscriber_status": {
        "computed": true,
        "description_kind": "plain",
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
      "source": {
        "block": {
          "block_types": {
            "aws_log_source_resource": {
              "block": {
                "attributes": {
                  "source_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "custom_log_source_resource": {
              "block": {
                "attributes": {
                  "attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "crawler_arn": "string",
                          "database_arn": "string",
                          "table_arn": "string"
                        }
                      ]
                    ]
                  },
                  "provider": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "location": "string",
                          "role_arn": "string"
                        }
                      ]
                    ]
                  },
                  "source_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_version": {
                    "computed": true,
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
        "nesting_mode": "set"
      },
      "subscriber_identity": {
        "block": {
          "attributes": {
            "external_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "principal": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsSecuritylakeSubscriberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecuritylakeSubscriber), &result)
	return &result
}
