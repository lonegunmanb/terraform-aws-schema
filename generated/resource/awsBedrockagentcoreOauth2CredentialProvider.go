package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreOauth2CredentialProvider = `{
  "block": {
    "attributes": {
      "client_secret_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "secret_arn": "string"
            }
          ]
        ]
      },
      "credential_provider_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_vendor": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      }
    },
    "block_types": {
      "oauth2_provider_config": {
        "block": {
          "block_types": {
            "custom_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  }
                },
                "block_types": {
                  "oauth_discovery": {
                    "block": {
                      "attributes": {
                        "discovery_url": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "authorization_server_metadata": {
                          "block": {
                            "attributes": {
                              "authorization_endpoint": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "issuer": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "response_types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "token_endpoint": {
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
            },
            "github_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_server_metadata": [
                            "list",
                            [
                              "object",
                              {
                                "authorization_endpoint": "string",
                                "issuer": "string",
                                "response_types": [
                                  "set",
                                  "string"
                                ],
                                "token_endpoint": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "google_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_server_metadata": [
                            "list",
                            [
                              "object",
                              {
                                "authorization_endpoint": "string",
                                "issuer": "string",
                                "response_types": [
                                  "set",
                                  "string"
                                ],
                                "token_endpoint": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "microsoft_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_server_metadata": [
                            "list",
                            [
                              "object",
                              {
                                "authorization_endpoint": "string",
                                "issuer": "string",
                                "response_types": [
                                  "set",
                                  "string"
                                ],
                                "token_endpoint": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "salesforce_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_server_metadata": [
                            "list",
                            [
                              "object",
                              {
                                "authorization_endpoint": "string",
                                "issuer": "string",
                                "response_types": [
                                  "set",
                                  "string"
                                ],
                                "token_endpoint": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "slack_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_credentials_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_server_metadata": [
                            "list",
                            [
                              "object",
                              {
                                "authorization_endpoint": "string",
                                "issuer": "string",
                                "response_types": [
                                  "set",
                                  "string"
                                ],
                                "token_endpoint": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
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

func AwsBedrockagentcoreOauth2CredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreOauth2CredentialProvider), &result)
	return &result
}
