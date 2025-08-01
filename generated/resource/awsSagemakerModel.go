package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerModel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_network_isolation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      "container": {
        "block": {
          "attributes": {
            "container_hostname": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inference_specification_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_data_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_package_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "image_config": {
              "block": {
                "attributes": {
                  "repository_access_mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "repository_auth_config": {
                    "block": {
                      "attributes": {
                        "repository_credentials_provider_arn": {
                          "description_kind": "plain",
                          "required": true,
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
            },
            "model_data_source": {
              "block": {
                "block_types": {
                  "s3_data_source": {
                    "block": {
                      "attributes": {
                        "compression_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_data_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "model_access_config": {
                          "block": {
                            "attributes": {
                              "accept_eula": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "multi_model_config": {
              "block": {
                "attributes": {
                  "model_cache_setting": {
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
        "nesting_mode": "list"
      },
      "inference_execution_config": {
        "block": {
          "attributes": {
            "mode": {
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
      "primary_container": {
        "block": {
          "attributes": {
            "container_hostname": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "image": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inference_specification_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_data_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_package_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "image_config": {
              "block": {
                "attributes": {
                  "repository_access_mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "repository_auth_config": {
                    "block": {
                      "attributes": {
                        "repository_credentials_provider_arn": {
                          "description_kind": "plain",
                          "required": true,
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
            },
            "model_data_source": {
              "block": {
                "block_types": {
                  "s3_data_source": {
                    "block": {
                      "attributes": {
                        "compression_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_data_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "model_access_config": {
                          "block": {
                            "attributes": {
                              "accept_eula": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "multi_model_config": {
              "block": {
                "attributes": {
                  "model_cache_setting": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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

func AwsSagemakerModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerModel), &result)
	return &result
}
