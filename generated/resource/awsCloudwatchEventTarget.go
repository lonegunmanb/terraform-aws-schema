package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchEventTarget = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_bus_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_path": {
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
      "role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "appsync_target": {
        "block": {
          "attributes": {
            "graphql_operation": {
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
      "batch_target": {
        "block": {
          "attributes": {
            "array_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "job_attempts": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "job_definition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "job_name": {
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
      "dead_letter_config": {
        "block": {
          "attributes": {
            "arn": {
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
      "ecs_target": {
        "block": {
          "attributes": {
            "enable_ecs_managed_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_execute_command": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "launch_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "platform_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "propagate_tags": {
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
            "task_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_definition_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "capacity_provider_strategy": {
              "block": {
                "attributes": {
                  "base": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "capacity_provider": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "weight": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "network_configuration": {
              "block": {
                "attributes": {
                  "assign_public_ip": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "security_groups": {
                    "description_kind": "plain",
                    "optional": true,
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
            },
            "ordered_placement_strategy": {
              "block": {
                "attributes": {
                  "field": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "placement_constraint": {
              "block": {
                "attributes": {
                  "expression": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 10,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_target": {
        "block": {
          "attributes": {
            "header_parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "path_parameter_values": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_string_parameters": {
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
      "input_transformer": {
        "block": {
          "attributes": {
            "input_paths": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "input_template": {
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
      "kinesis_target": {
        "block": {
          "attributes": {
            "partition_key_path": {
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
      "redshift_target": {
        "block": {
          "attributes": {
            "database": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "db_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "statement_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "with_event": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retry_policy": {
        "block": {
          "attributes": {
            "maximum_event_age_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maximum_retry_attempts": {
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
      "run_command_targets": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      },
      "sagemaker_pipeline_target": {
        "block": {
          "block_types": {
            "pipeline_parameter_list": {
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
              "max_items": 200,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sqs_target": {
        "block": {
          "attributes": {
            "message_group_id": {
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
  "version": 1
}`

func AwsCloudwatchEventTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchEventTarget), &result)
	return &result
}
