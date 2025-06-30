package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerSpace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_efs_file_system_uid": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "space_display_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "space_name": {
        "description_kind": "plain",
        "required": true,
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
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "ownership_settings": {
        "block": {
          "attributes": {
            "owner_user_profile_name": {
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
      "space_settings": {
        "block": {
          "attributes": {
            "app_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "code_editor_app_settings": {
              "block": {
                "block_types": {
                  "app_lifecycle_management": {
                    "block": {
                      "block_types": {
                        "idle_settings": {
                          "block": {
                            "attributes": {
                              "idle_timeout_in_minutes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
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
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_alias": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
                          "description_kind": "plain",
                          "optional": true,
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
            },
            "custom_file_system": {
              "block": {
                "block_types": {
                  "efs_file_system": {
                    "block": {
                      "attributes": {
                        "file_system_id": {
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
              "nesting_mode": "list"
            },
            "jupyter_lab_app_settings": {
              "block": {
                "block_types": {
                  "app_lifecycle_management": {
                    "block": {
                      "block_types": {
                        "idle_settings": {
                          "block": {
                            "attributes": {
                              "idle_timeout_in_minutes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
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
                  "code_repository": {
                    "block": {
                      "attributes": {
                        "repository_url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "nesting_mode": "set"
                  },
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_alias": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
                          "description_kind": "plain",
                          "optional": true,
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
            },
            "jupyter_server_app_settings": {
              "block": {
                "attributes": {
                  "lifecycle_config_arns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "code_repository": {
                    "block": {
                      "attributes": {
                        "repository_url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "nesting_mode": "set"
                  },
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_alias": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
                          "description_kind": "plain",
                          "optional": true,
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
            },
            "kernel_gateway_app_settings": {
              "block": {
                "attributes": {
                  "lifecycle_config_arns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "custom_image": {
                    "block": {
                      "attributes": {
                        "app_image_config_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_version_number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 200,
                    "nesting_mode": "list"
                  },
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_alias": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
                          "description_kind": "plain",
                          "optional": true,
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
            },
            "space_storage_settings": {
              "block": {
                "block_types": {
                  "ebs_storage_settings": {
                    "block": {
                      "attributes": {
                        "ebs_volume_size_in_gb": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "space_sharing_settings": {
        "block": {
          "attributes": {
            "sharing_type": {
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
  "version": 0
}`

func AwsSagemakerSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerSpace), &result)
	return &result
}
