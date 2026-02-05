package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsArcregionswitchPlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recovery_approach": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_time_objective_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "deprecated": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
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
      "associated_alarms": {
        "block": {
          "attributes": {
            "alarm_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cross_account_role": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "map_block_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      },
      "triggers": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_delay_minutes_between_executions": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "target_region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "conditions": {
              "block": {
                "attributes": {
                  "associated_alarm_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "condition": {
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
      "workflow": {
        "block": {
          "attributes": {
            "workflow_description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "workflow_target_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workflow_target_region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "step": {
              "block": {
                "attributes": {
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "execution_block_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "arc_routing_control_config": {
                    "block": {
                      "attributes": {
                        "cross_account_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "external_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "region_and_routing_controls": {
                          "block": {
                            "attributes": {
                              "region": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "routing_control": {
                                "block": {
                                  "attributes": {
                                    "routing_control_arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "state": {
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
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "custom_action_lambda_config": {
                    "block": {
                      "attributes": {
                        "region_to_run": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "retry_interval_minutes": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "lambda": {
                          "block": {
                            "attributes": {
                              "arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "cross_account_role": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "external_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "behavior": {
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
                  "document_db_config": {
                    "block": {
                      "attributes": {
                        "behavior": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "cross_account_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_cluster_arns": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "external_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "global_cluster_identifier": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "ungraceful": {
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
                  "ec2_asg_capacity_increase_config": {
                    "block": {
                      "attributes": {
                        "capacity_monitoring_approach": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_percent": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "asg": {
                          "block": {
                            "attributes": {
                              "arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "cross_account_role": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "external_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "minimum_success_percentage": {
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
                  "ecs_capacity_increase_config": {
                    "block": {
                      "attributes": {
                        "capacity_monitoring_approach": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_percent": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "service": {
                          "block": {
                            "attributes": {
                              "cluster_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "cross_account_role": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "external_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "service_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "minimum_success_percentage": {
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
                  "eks_resource_scaling_config": {
                    "block": {
                      "attributes": {
                        "capacity_monitoring_approach": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_percent": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "eks_clusters": {
                          "block": {
                            "attributes": {
                              "cluster_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "cross_account_role": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "external_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "kubernetes_resource_type": {
                          "block": {
                            "attributes": {
                              "api_version": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "kind": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "scaling_resources": {
                          "block": {
                            "attributes": {
                              "namespace": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "resources": {
                                "block": {
                                  "attributes": {
                                    "hpa_name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "resource_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "set"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "minimum_success_percentage": {
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
                  "execution_approval_config": {
                    "block": {
                      "attributes": {
                        "approval_role": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "global_aurora_config": {
                    "block": {
                      "attributes": {
                        "behavior": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "cross_account_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_cluster_arns": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "external_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "global_cluster_identifier": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "ungraceful": {
                          "block": {
                            "attributes": {
                              "ungraceful": {
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
                  "parallel_config": {
                    "block": {
                      "block_types": {
                        "step": {
                          "block": {
                            "attributes": {
                              "description": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "execution_block_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "arc_routing_control_config": {
                                "block": {
                                  "attributes": {
                                    "cross_account_role": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "external_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "region_and_routing_controls": {
                                      "block": {
                                        "attributes": {
                                          "region": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "routing_control": {
                                            "block": {
                                              "attributes": {
                                                "routing_control_arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "state": {
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
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "custom_action_lambda_config": {
                                "block": {
                                  "attributes": {
                                    "region_to_run": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "retry_interval_minutes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "lambda": {
                                      "block": {
                                        "attributes": {
                                          "arn": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "cross_account_role": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "external_id": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "behavior": {
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
                              "document_db_config": {
                                "block": {
                                  "attributes": {
                                    "behavior": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "cross_account_role": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "database_cluster_arns": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "external_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "global_cluster_identifier": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "ungraceful": {
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
                              "ec2_asg_capacity_increase_config": {
                                "block": {
                                  "attributes": {
                                    "capacity_monitoring_approach": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "target_percent": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "asg": {
                                      "block": {
                                        "attributes": {
                                          "arn": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "cross_account_role": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "external_id": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "minimum_success_percentage": {
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
                              "ecs_capacity_increase_config": {
                                "block": {
                                  "attributes": {
                                    "capacity_monitoring_approach": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "target_percent": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "service": {
                                      "block": {
                                        "attributes": {
                                          "cluster_arn": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "cross_account_role": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "external_id": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "service_arn": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "minimum_success_percentage": {
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
                              "eks_resource_scaling_config": {
                                "block": {
                                  "attributes": {
                                    "capacity_monitoring_approach": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "target_percent": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "eks_clusters": {
                                      "block": {
                                        "attributes": {
                                          "cluster_arn": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "cross_account_role": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "external_id": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "kubernetes_resource_type": {
                                      "block": {
                                        "attributes": {
                                          "api_version": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "kind": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "scaling_resources": {
                                      "block": {
                                        "attributes": {
                                          "namespace": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "resources": {
                                            "block": {
                                              "attributes": {
                                                "hpa_name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "namespace": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "resource_name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "minimum_success_percentage": {
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
                              "execution_approval_config": {
                                "block": {
                                  "attributes": {
                                    "approval_role": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "global_aurora_config": {
                                "block": {
                                  "attributes": {
                                    "behavior": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "cross_account_role": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "database_cluster_arns": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "external_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "global_cluster_identifier": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "ungraceful": {
                                      "block": {
                                        "attributes": {
                                          "ungraceful": {
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
                              "region_switch_plan_config": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "cross_account_role": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "external_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "route53_health_check_config": {
                                "block": {
                                  "attributes": {
                                    "cross_account_role": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "external_id": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "hosted_zone_id": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "record_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "timeout_minutes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "record_set": {
                                      "block": {
                                        "attributes": {
                                          "record_set_identifier": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "region": {
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
                  "region_switch_plan_config": {
                    "block": {
                      "attributes": {
                        "arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "cross_account_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "external_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "route53_health_check_config": {
                    "block": {
                      "attributes": {
                        "cross_account_role": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "external_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "hosted_zone_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "record_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "record_set": {
                          "block": {
                            "attributes": {
                              "record_set_identifier": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "region": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsArcregionswitchPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsArcregionswitchPlan), &result)
	return &result
}
