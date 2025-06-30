package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCustomerprofilesDomain = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_queue_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_encryption_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_expiration_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "domain_name": {
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
      "matching": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "auto_merging": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "min_allowed_confidence_score_for_merging": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "conflict_resolution": {
                    "block": {
                      "attributes": {
                        "conflict_resolving_model": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "source_name": {
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
                  "consolidation": {
                    "block": {
                      "attributes": {
                        "matching_attributes_list": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            [
                              "list",
                              "string"
                            ]
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "exporting_config": {
              "block": {
                "block_types": {
                  "s3_exporting": {
                    "block": {
                      "attributes": {
                        "s3_bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_key_name": {
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
            "job_schedule": {
              "block": {
                "attributes": {
                  "day_of_the_week": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "time": {
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
      "rule_based_matching": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "max_allowed_rule_level_for_matching": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_allowed_rule_level_for_merging": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "attribute_types_selector": {
              "block": {
                "attributes": {
                  "address": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "attribute_matching_model": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "email_address": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "phone_number": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "conflict_resolution": {
              "block": {
                "attributes": {
                  "conflict_resolving_model": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_name": {
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
            "exporting_config": {
              "block": {
                "block_types": {
                  "s3_exporting": {
                    "block": {
                      "attributes": {
                        "s3_bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_key_name": {
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
            "matching_rules": {
              "block": {
                "attributes": {
                  "rule": {
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
              "nesting_mode": "set"
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

func AwsCustomerprofilesDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCustomerprofilesDomain), &result)
	return &result
}
