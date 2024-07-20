package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskReplicator = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "current_version": {
        "computed": true,
        "description_kind": "plain",
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
      "replicator_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_execution_role_arn": {
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
      }
    },
    "block_types": {
      "kafka_cluster": {
        "block": {
          "block_types": {
            "amazon_msk_cluster": {
              "block": {
                "attributes": {
                  "msk_cluster_arn": {
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
            },
            "vpc_config": {
              "block": {
                "attributes": {
                  "security_groups_ids": {
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
        "max_items": 2,
        "min_items": 2,
        "nesting_mode": "list"
      },
      "replication_info_list": {
        "block": {
          "attributes": {
            "source_kafka_cluster_alias": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_kafka_cluster_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_compression_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_kafka_cluster_alias": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_kafka_cluster_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "consumer_group_replication": {
              "block": {
                "attributes": {
                  "consumer_groups_to_exclude": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "consumer_groups_to_replicate": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "detect_and_copy_new_consumer_groups": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "synchronise_consumer_group_offsets": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "topic_replication": {
              "block": {
                "attributes": {
                  "copy_access_control_lists_for_topics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "copy_topic_configurations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "detect_and_copy_new_topics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "topics_to_exclude": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "topics_to_replicate": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "starting_position": {
                    "block": {
                      "attributes": {
                        "type": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsMskReplicatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskReplicator), &result)
	return &result
}
