package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3Bucket = `{
  "block": {
    "attributes": {
      "acceleration_status": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket_regional_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "force_destroy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hosted_zone_id": {
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
      "object_lock_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "policy": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "request_payer": {
        "computed": true,
        "deprecated": true,
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
      },
      "website_domain": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "website_endpoint": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cors_rule": {
        "block": {
          "attributes": {
            "allowed_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allowed_origins": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "expose_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_age_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "grant": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "permissions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "lifecycle_rule": {
        "block": {
          "attributes": {
            "abort_incomplete_multipart_upload_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix": {
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
            }
          },
          "block_types": {
            "expiration": {
              "block": {
                "attributes": {
                  "date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "expired_object_delete_marker": {
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
            "noncurrent_version_expiration": {
              "block": {
                "attributes": {
                  "days": {
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
            "noncurrent_version_transition": {
              "block": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "transition": {
              "block": {
                "attributes": {
                  "date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
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
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "logging": {
        "block": {
          "attributes": {
            "target_bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "object_lock_configuration": {
        "block": {
          "attributes": {
            "object_lock_enabled": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "rule": {
              "block": {
                "block_types": {
                  "default_retention": {
                    "block": {
                      "attributes": {
                        "days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "years": {
                          "description_kind": "plain",
                          "optional": true,
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
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "replication_configuration": {
        "block": {
          "attributes": {
            "role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "delete_marker_replication_status": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "destination": {
                    "block": {
                      "attributes": {
                        "account_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "replica_kms_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "storage_class": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "access_control_translation": {
                          "block": {
                            "attributes": {
                              "owner": {
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
                        "metrics": {
                          "block": {
                            "attributes": {
                              "minutes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "status": {
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
                        "replication_time": {
                          "block": {
                            "attributes": {
                              "minutes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "status": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter": {
                    "block": {
                      "attributes": {
                        "prefix": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "source_selection_criteria": {
                    "block": {
                      "block_types": {
                        "sse_kms_encrypted_objects": {
                          "block": {
                            "attributes": {
                              "enabled": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "server_side_encryption_configuration": {
        "block": {
          "block_types": {
            "rule": {
              "block": {
                "attributes": {
                  "bucket_key_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "apply_server_side_encryption_by_default": {
                    "block": {
                      "attributes": {
                        "kms_master_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sse_algorithm": {
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
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
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
            "read": {
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
      },
      "versioning": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "mfa_delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "website": {
        "block": {
          "attributes": {
            "error_document": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_document": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_all_requests_to": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "routing_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
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

func AwsS3BucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3Bucket), &result)
	return &result
}
