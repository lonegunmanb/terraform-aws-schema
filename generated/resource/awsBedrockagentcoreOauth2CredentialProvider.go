package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockagentcoreOauth2CredentialProvider = `{
  "block": {
    "attributes": {
      "callback_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
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
      "oauth2_provider_config": {
        "block": {
          "block_types": {
            "atlassian_oauth2_provider_config": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
            "custom_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "client_authentication_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
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
                              },
                              "token_endpoint_auth_methods": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
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
                  },
                  "on_behalf_of_token_exchange_config": {
                    "block": {
                      "attributes": {
                        "grant_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "token_exchange_grant_type_config": {
                          "block": {
                            "attributes": {
                              "actor_token_content": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "actor_token_scopes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
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
                  },
                  "private_endpoint": {
                    "block": {
                      "block_types": {
                        "managed_vpc_resource": {
                          "block": {
                            "attributes": {
                              "endpoint_ip_address_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "routing_domain": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "security_group_ids": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "subnet_ids": {
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "tags": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "vpc_identifier": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "self_managed_lattice_resource": {
                          "block": {
                            "attributes": {
                              "resource_configuration_identifier": {
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
                  },
                  "private_endpoint_override": {
                    "block": {
                      "attributes": {
                        "domain": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "private_endpoint": {
                          "block": {
                            "block_types": {
                              "managed_vpc_resource": {
                                "block": {
                                  "attributes": {
                                    "endpoint_ip_address_type": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "routing_domain": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "security_group_ids": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "subnet_ids": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "tags": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "vpc_identifier": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "self_managed_lattice_resource": {
                                "block": {
                                  "attributes": {
                                    "resource_configuration_identifier": {
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
                  },
                  "private_key_jwt_config": {
                    "block": {
                      "attributes": {
                        "additional_header_claims": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "additional_payload_claims": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "signing_algorithm": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "private_key_source": {
                          "block": {
                            "block_types": {
                              "kms_key_source": {
                                "block": {
                                  "attributes": {
                                    "kms_key_arn": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
            "included_oauth2_provider_config": {
              "block": {
                "attributes": {
                  "authorization_endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_secret_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "issuer": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  },
                  "token_endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
            "linkedin_oauth2_provider_config": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  },
                  "tenant_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "tenant_id_wo": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string",
                    "write_only": true
                  },
                  "tenant_id_wo_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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
                  "client_secret_source": {
                    "description_kind": "plain",
                    "optional": true,
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
                                "token_endpoint": "string",
                                "token_endpoint_auth_methods": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "discovery_url": "string"
                        }
                      ]
                    ]
                  }
                },
                "block_types": {
                  "client_secret_config": {
                    "block": {
                      "attributes": {
                        "json_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_id": {
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

func AwsBedrockagentcoreOauth2CredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockagentcoreOauth2CredentialProvider), &result)
	return &result
}
