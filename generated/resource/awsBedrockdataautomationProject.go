package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockdataautomationProject = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_stage": {
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
      "type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "custom_output_configuration": {
        "block": {
          "block_types": {
            "blueprint": {
              "block": {
                "attributes": {
                  "blueprint_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "blueprint_stage": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "blueprint_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "document": {
              "block": {
                "block_types": {
                  "fallback_blueprint": {
                    "block": {
                      "attributes": {
                        "blueprint_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "blueprint_stage": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "blueprint_version": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "data_automation_library_configuration": {
        "block": {
          "block_types": {
            "library": {
              "block": {
                "attributes": {
                  "library_arn": {
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
      "encryption_configuration": {
        "block": {
          "attributes": {
            "kms_encryption_context": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "kms_key_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "override_configuration": {
        "block": {
          "block_types": {
            "audio": {
              "block": {
                "block_types": {
                  "language_configuration": {
                    "block": {
                      "attributes": {
                        "generative_output_language": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "identify_multiple_languages": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "input_languages": {
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
                  "modality_processing": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sensitive_data_configuration": {
                    "block": {
                      "attributes": {
                        "detection_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "detection_scope": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "pii_entities_configuration": {
                          "block": {
                            "attributes": {
                              "pii_entity_types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "redaction_mask_mode": {
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "document": {
              "block": {
                "block_types": {
                  "modality_processing": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sensitive_data_configuration": {
                    "block": {
                      "attributes": {
                        "detection_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "detection_scope": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "pii_entities_configuration": {
                          "block": {
                            "attributes": {
                              "pii_entity_types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "redaction_mask_mode": {
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
                  "splitter": {
                    "block": {
                      "attributes": {
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
              "nesting_mode": "list"
            },
            "image": {
              "block": {
                "block_types": {
                  "modality_processing": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sensitive_data_configuration": {
                    "block": {
                      "attributes": {
                        "detection_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "detection_scope": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "pii_entities_configuration": {
                          "block": {
                            "attributes": {
                              "pii_entity_types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "redaction_mask_mode": {
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "modality_routing": {
              "block": {
                "attributes": {
                  "jpeg": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mov": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mp4": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "png": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "video": {
              "block": {
                "block_types": {
                  "modality_processing": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "sensitive_data_configuration": {
                    "block": {
                      "attributes": {
                        "detection_mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "detection_scope": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "pii_entities_configuration": {
                          "block": {
                            "attributes": {
                              "pii_entity_types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "redaction_mask_mode": {
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
      "standard_output_configuration": {
        "block": {
          "block_types": {
            "audio": {
              "block": {
                "block_types": {
                  "extraction": {
                    "block": {
                      "block_types": {
                        "category": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "types": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "type_configuration": {
                                "block": {
                                  "block_types": {
                                    "transcript": {
                                      "block": {
                                        "block_types": {
                                          "channel_labeling": {
                                            "block": {
                                              "attributes": {
                                                "state": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "speaker_labeling": {
                                            "block": {
                                              "attributes": {
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
                  "generative_field": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "types": {
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "document": {
              "block": {
                "block_types": {
                  "extraction": {
                    "block": {
                      "block_types": {
                        "bounding_box": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "granularity": {
                          "block": {
                            "attributes": {
                              "types": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "generative_field": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "output_format": {
                    "block": {
                      "block_types": {
                        "additional_file_format": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "text_format": {
                          "block": {
                            "attributes": {
                              "types": {
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
            "image": {
              "block": {
                "block_types": {
                  "extraction": {
                    "block": {
                      "block_types": {
                        "bounding_box": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "category": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "types": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "generative_field": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "types": {
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "video": {
              "block": {
                "block_types": {
                  "extraction": {
                    "block": {
                      "block_types": {
                        "bounding_box": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "category": {
                          "block": {
                            "attributes": {
                              "state": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "types": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "generative_field": {
                    "block": {
                      "attributes": {
                        "state": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "types": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockdataautomationProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockdataautomationProject), &result)
	return &result
}
