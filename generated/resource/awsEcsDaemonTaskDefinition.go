package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsDaemonTaskDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "family": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "memory": {
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
      "revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
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
      },
      "task_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "container_definition": {
        "block": {
          "attributes": {
            "command": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "entry_point": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "essential": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "interactive": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "memory": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory_reservation": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "privileged": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pseudo_terminal": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "readonly_root_filesystem": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "start_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stop_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "user": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "working_directory": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "depends_on": {
              "block": {
                "attributes": {
                  "condition": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "container_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "environment": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "environment_file": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "firelens_configuration": {
              "block": {
                "attributes": {
                  "options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "health_check": {
              "block": {
                "attributes": {
                  "command": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retries": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "linux_parameters": {
              "block": {
                "attributes": {
                  "init_process_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "capabilities": {
                    "block": {
                      "attributes": {
                        "add": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "drop": {
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
                  },
                  "device": {
                    "block": {
                      "attributes": {
                        "container_path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "permissions": {
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
                  },
                  "tmpfs": {
                    "block": {
                      "attributes": {
                        "container_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "mount_options": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
            "log_configuration": {
              "block": {
                "attributes": {
                  "log_driver": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "secret_option": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_from": {
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
            "mount_point": {
              "block": {
                "attributes": {
                  "container_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_volume": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "repository_credentials": {
              "block": {
                "attributes": {
                  "credentials_parameter": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "restart_policy": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "ignored_exit_codes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "restart_attempt_period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "secret": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value_from": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "system_control": {
              "block": {
                "attributes": {
                  "namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ulimit": {
              "block": {
                "attributes": {
                  "hard_limit": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "soft_limit": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
      "volume": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "host": {
              "block": {
                "attributes": {
                  "source_path": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsDaemonTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsDaemonTaskDefinition), &result)
	return &result
}
