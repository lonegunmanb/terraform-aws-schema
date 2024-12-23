package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBatchJobDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deregister_on_new_revision": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ecs_properties": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "platform_capabilities": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "propagate_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scheduling_priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "eks_properties": {
        "block": {
          "block_types": {
            "pod_properties": {
              "block": {
                "attributes": {
                  "dns_policy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host_network": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "service_account_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "share_process_namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "containers": {
                    "block": {
                      "attributes": {
                        "args": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "image": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_pull_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "env": {
                          "block": {
                            "attributes": {
                              "name": {
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
                          "nesting_mode": "set"
                        },
                        "resources": {
                          "block": {
                            "attributes": {
                              "limits": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "requests": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "security_context": {
                          "block": {
                            "attributes": {
                              "privileged": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "read_only_root_file_system": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "run_as_group": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "run_as_non_root": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "run_as_user": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "volume_mounts": {
                          "block": {
                            "attributes": {
                              "mount_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "read_only": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "image_pull_secret": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "init_containers": {
                    "block": {
                      "attributes": {
                        "args": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "image": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_pull_policy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "env": {
                          "block": {
                            "attributes": {
                              "name": {
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
                          "nesting_mode": "set"
                        },
                        "resources": {
                          "block": {
                            "attributes": {
                              "limits": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "requests": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "security_context": {
                          "block": {
                            "attributes": {
                              "privileged": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "read_only_root_file_system": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "run_as_group": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "run_as_non_root": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "run_as_user": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "volume_mounts": {
                          "block": {
                            "attributes": {
                              "mount_path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "read_only": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "nesting_mode": "list"
                  },
                  "metadata": {
                    "block": {
                      "attributes": {
                        "labels": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "volumes": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "empty_dir": {
                          "block": {
                            "attributes": {
                              "medium": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "size_limit": {
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
                        "host_path": {
                          "block": {
                            "attributes": {
                              "path": {
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
                        "secret": {
                          "block": {
                            "attributes": {
                              "optional": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "secret_name": {
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
                    "nesting_mode": "list"
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
      "retry_strategy": {
        "block": {
          "attributes": {
            "attempts": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "evaluate_on_exit": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "on_exit_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_reason": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_status_reason": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeout": {
        "block": {
          "attributes": {
            "attempt_duration_seconds": {
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
  "version": 0
}`

func AwsBatchJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBatchJobDefinition), &result)
	return &result
}
