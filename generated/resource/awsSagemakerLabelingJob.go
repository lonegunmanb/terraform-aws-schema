package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerLabelingJob = `{
  "block": {
    "attributes": {
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_reference_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "label_attribute_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "label_category_config_s3_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_counters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "failed_non_retryable_error": "number",
              "human_labeled": "number",
              "machine_labeled": "number",
              "total_labeled": "number",
              "unlabeled": "number"
            }
          ]
        ]
      },
      "labeling_job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "labeling_job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labeling_job_status": {
        "computed": true,
        "description_kind": "plain",
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
        "required": true,
        "type": "string"
      },
      "stopping_conditions": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "max_human_labeled_object_count": "number",
              "max_percentage_of_input_dataset_labeled": "number"
            }
          ]
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
      "human_task_config": {
        "block": {
          "attributes": {
            "max_concurrent_task_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "number_of_human_workers_per_data_object": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "pre_human_task_lambda_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "task_availability_lifetime_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_description": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "task_keywords": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "task_time_limit_in_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "task_title": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workteam_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "annotation_consolidation_config": {
              "block": {
                "attributes": {
                  "annotation_consolidation_lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "public_workforce_task_price": {
              "block": {
                "block_types": {
                  "amount_in_usd": {
                    "block": {
                      "attributes": {
                        "cents": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "dollars": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tenth_fractions_of_a_cent": {
                          "description_kind": "plain",
                          "optional": true,
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
            "ui_config": {
              "block": {
                "attributes": {
                  "human_task_ui_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ui_template_s3_uri": {
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
      "input_config": {
        "block": {
          "block_types": {
            "data_attributes": {
              "block": {
                "attributes": {
                  "content_classifiers": {
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
            },
            "data_source": {
              "block": {
                "block_types": {
                  "s3_data_source": {
                    "block": {
                      "attributes": {
                        "manifest_s3_uri": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sns_data_source": {
                    "block": {
                      "attributes": {
                        "sns_topic_arn": {
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
      "labeling_job_algorithms_config": {
        "block": {
          "attributes": {
            "initial_active_learning_model_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labeling_job_algorithm_specification_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "labeling_job_resource_config": {
              "block": {
                "attributes": {
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
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
      "output_config": {
        "block": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_output_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sns_topic_arn": {
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
  "version": 0
}`

func AwsSagemakerLabelingJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerLabelingJob), &result)
	return &result
}
