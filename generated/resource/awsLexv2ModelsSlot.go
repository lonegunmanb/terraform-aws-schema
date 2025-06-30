package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLexv2ModelsSlot = `{
  "block": {
    "attributes": {
      "bot_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bot_version": {
        "description_kind": "plain",
        "required": true,
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
        "type": "string"
      },
      "intent_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locale_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slot_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "slot_type_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "multiple_values_setting": {
        "block": {
          "attributes": {
            "allow_multiple_values": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "obfuscation_setting": {
        "block": {
          "attributes": {
            "obfuscation_setting_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sub_slot_setting": {
        "block": {
          "attributes": {
            "expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "slot_specification": {
              "block": {
                "attributes": {
                  "map_block_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "slot_type_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "value_elicitation_setting": {
                    "block": {
                      "block_types": {
                        "default_value_specification": {
                          "block": {
                            "block_types": {
                              "default_value_list": {
                                "block": {
                                  "attributes": {
                                    "default_value": {
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
                        "prompt_specification": {
                          "block": {
                            "attributes": {
                              "allow_interrupt": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "max_retries": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "message_selection_strategy": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "message_group": {
                                "block": {
                                  "block_types": {
                                    "message": {
                                      "block": {
                                        "block_types": {
                                          "custom_payload": {
                                            "block": {
                                              "attributes": {
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
                                          "image_response_card": {
                                            "block": {
                                              "attributes": {
                                                "image_url": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "button": {
                                                  "block": {
                                                    "attributes": {
                                                      "text": {
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
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "plain_text_message": {
                                            "block": {
                                              "attributes": {
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
                                          "ssml_message": {
                                            "block": {
                                              "attributes": {
                                                "value": {
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
                                    "variation": {
                                      "block": {
                                        "block_types": {
                                          "custom_payload": {
                                            "block": {
                                              "attributes": {
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
                                          "image_response_card": {
                                            "block": {
                                              "attributes": {
                                                "image_url": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "button": {
                                                  "block": {
                                                    "attributes": {
                                                      "text": {
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
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "plain_text_message": {
                                            "block": {
                                              "attributes": {
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
                                          "ssml_message": {
                                            "block": {
                                              "attributes": {
                                                "value": {
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
                              "prompt_attempts_specification": {
                                "block": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "map_block_key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "allowed_input_types": {
                                      "block": {
                                        "attributes": {
                                          "allow_audio_input": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          },
                                          "allow_dtmf_input": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "bool"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "audio_and_dtmf_input_specification": {
                                      "block": {
                                        "attributes": {
                                          "start_timeout_ms": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          }
                                        },
                                        "block_types": {
                                          "audio_specification": {
                                            "block": {
                                              "attributes": {
                                                "end_timeout_ms": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "max_length_ms": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          },
                                          "dtmf_specification": {
                                            "block": {
                                              "attributes": {
                                                "deletion_character": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "end_character": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "end_timeout_ms": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "max_length": {
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
                                    "text_input_specification": {
                                      "block": {
                                        "attributes": {
                                          "start_timeout_ms": {
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
                                "nesting_mode": "set"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "sample_utterance": {
                          "block": {
                            "attributes": {
                              "utterance": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "wait_and_continue_specification": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "continue_response": {
                                "block": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "message_group": {
                                      "block": {
                                        "block_types": {
                                          "message": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
                                          "variation": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
                              "still_waiting_response": {
                                "block": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "frequency_in_seconds": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "timeout_in_seconds": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "message_group": {
                                      "block": {
                                        "block_types": {
                                          "message": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
                                          "variation": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
                              "waiting_response": {
                                "block": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "message_group": {
                                      "block": {
                                        "block_types": {
                                          "message": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
                                          "variation": {
                                            "block": {
                                              "block_types": {
                                                "custom_payload": {
                                                  "block": {
                                                    "attributes": {
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
                                                "image_response_card": {
                                                  "block": {
                                                    "attributes": {
                                                      "image_url": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "button": {
                                                        "block": {
                                                          "attributes": {
                                                            "text": {
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
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "nesting_mode": "list"
                                                },
                                                "plain_text_message": {
                                                  "block": {
                                                    "attributes": {
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
                                                "ssml_message": {
                                                  "block": {
                                                    "attributes": {
                                                      "value": {
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
      "value_elicitation_setting": {
        "block": {
          "attributes": {
            "slot_constraint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_value_specification": {
              "block": {
                "block_types": {
                  "default_value_list": {
                    "block": {
                      "attributes": {
                        "default_value": {
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
            "prompt_specification": {
              "block": {
                "attributes": {
                  "allow_interrupt": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "max_retries": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "message_selection_strategy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "message_group": {
                    "block": {
                      "block_types": {
                        "message": {
                          "block": {
                            "block_types": {
                              "custom_payload": {
                                "block": {
                                  "attributes": {
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
                              "image_response_card": {
                                "block": {
                                  "attributes": {
                                    "image_url": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "subtitle": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "title": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "button": {
                                      "block": {
                                        "attributes": {
                                          "text": {
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
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "plain_text_message": {
                                "block": {
                                  "attributes": {
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
                              "ssml_message": {
                                "block": {
                                  "attributes": {
                                    "value": {
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
                        "variation": {
                          "block": {
                            "block_types": {
                              "custom_payload": {
                                "block": {
                                  "attributes": {
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
                              "image_response_card": {
                                "block": {
                                  "attributes": {
                                    "image_url": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "subtitle": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "title": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "button": {
                                      "block": {
                                        "attributes": {
                                          "text": {
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
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "plain_text_message": {
                                "block": {
                                  "attributes": {
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
                              "ssml_message": {
                                "block": {
                                  "attributes": {
                                    "value": {
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
                  "prompt_attempts_specification": {
                    "block": {
                      "attributes": {
                        "allow_interrupt": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "map_block_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "allowed_input_types": {
                          "block": {
                            "attributes": {
                              "allow_audio_input": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              },
                              "allow_dtmf_input": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "audio_and_dtmf_input_specification": {
                          "block": {
                            "attributes": {
                              "start_timeout_ms": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "audio_specification": {
                                "block": {
                                  "attributes": {
                                    "end_timeout_ms": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "max_length_ms": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "dtmf_specification": {
                                "block": {
                                  "attributes": {
                                    "deletion_character": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "end_character": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "end_timeout_ms": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "max_length": {
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
                        "text_input_specification": {
                          "block": {
                            "attributes": {
                              "start_timeout_ms": {
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
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "sample_utterance": {
              "block": {
                "attributes": {
                  "utterance": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "slot_resolution_setting": {
              "block": {
                "attributes": {
                  "slot_resolution_strategy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "wait_and_continue_specification": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "continue_response": {
                    "block": {
                      "attributes": {
                        "allow_interrupt": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "message_group": {
                          "block": {
                            "block_types": {
                              "message": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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
                              "variation": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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
                  "still_waiting_response": {
                    "block": {
                      "attributes": {
                        "allow_interrupt": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "frequency_in_seconds": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "timeout_in_seconds": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "message_group": {
                          "block": {
                            "block_types": {
                              "message": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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
                              "variation": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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
                  "waiting_response": {
                    "block": {
                      "attributes": {
                        "allow_interrupt": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "message_group": {
                          "block": {
                            "block_types": {
                              "message": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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
                              "variation": {
                                "block": {
                                  "block_types": {
                                    "custom_payload": {
                                      "block": {
                                        "attributes": {
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
                                    "image_response_card": {
                                      "block": {
                                        "attributes": {
                                          "image_url": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "subtitle": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "title": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "button": {
                                            "block": {
                                              "attributes": {
                                                "text": {
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
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "plain_text_message": {
                                      "block": {
                                        "attributes": {
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
                                    "ssml_message": {
                                      "block": {
                                        "attributes": {
                                          "value": {
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

func AwsLexv2ModelsSlotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLexv2ModelsSlot), &result)
	return &result
}
