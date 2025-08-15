package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppsyncApi = `{
  "block": {
    "attributes": {
      "api_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_contact": {
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
      "waf_web_acl_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "xray_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "block_types": {
      "event_config": {
        "block": {
          "block_types": {
            "auth_provider": {
              "block": {
                "attributes": {
                  "auth_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "cognito_config": {
                    "block": {
                      "attributes": {
                        "app_id_client_regex": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "aws_region": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user_pool_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "lambda_authorizer_config": {
                    "block": {
                      "attributes": {
                        "authorizer_result_ttl_in_seconds": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "authorizer_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "identity_validation_expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "openid_connect_config": {
                    "block": {
                      "attributes": {
                        "auth_ttl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "client_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iat_ttl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "issuer": {
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
            "connection_auth_mode": {
              "block": {
                "attributes": {
                  "auth_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "default_publish_auth_mode": {
              "block": {
                "attributes": {
                  "auth_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "default_subscribe_auth_mode": {
              "block": {
                "attributes": {
                  "auth_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "log_config": {
              "block": {
                "attributes": {
                  "cloudwatch_logs_role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "log_level": {
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
  "version": 0
}`

func AwsAppsyncApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppsyncApi), &result)
	return &result
}
