package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "athena_properties": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_properties": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "connection_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "match_criteria": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
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
      }
    },
    "block_types": {
      "authentication_configuration": {
        "block": {
          "attributes": {
            "authentication_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "custom_authentication_credentials": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "basic_authentication_credentials": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oauth2_properties": {
              "block": {
                "attributes": {
                  "oauth2_grant_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_url_parameters_map": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "authorization_code_properties": {
                    "block": {
                      "attributes": {
                        "authorization_code": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "redirect_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "oauth2_client_application": {
                    "block": {
                      "attributes": {
                        "aws_managed_client_application_reference": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_managed_client_application_client_id": {
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
                  "oauth2_credentials": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "jwt_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "refresh_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "user_managed_client_application_client_secret": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "physical_connection_requirements": {
        "block": {
          "attributes": {
            "availability_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_id_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
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
  "version": 0
}`

func AwsGlueConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueConnection), &result)
	return &result
}
