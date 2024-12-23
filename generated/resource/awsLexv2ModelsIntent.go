package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLexv2ModelsIntent = `{
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
      "creation_date_time": {
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
        "type": "string"
      },
      "intent_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_date_time": {
        "computed": true,
        "description_kind": "plain",
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
      "parent_intent_signature": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "closing_setting": {
        "block": {
          "attributes": {
            "active": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "closing_response": {
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
            "conditional": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "conditional_branch": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "expression_string": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
                  },
                  "default_branch": {
                    "block": {
                      "block_types": {
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
            },
            "next_step": {
              "block": {
                "attributes": {
                  "session_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dialog_action": {
                    "block": {
                      "attributes": {
                        "slot_to_elicit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suppress_next_message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "intent": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "slot": {
                          "block": {
                            "attributes": {
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "shape": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value": {
                                "block": {
                                  "attributes": {
                                    "interpreted_value": {
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
      "confirmation_setting": {
        "block": {
          "attributes": {
            "active": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "code_hook": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "enable_code_hook_invocation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "invocation_label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "post_code_hook_specification": {
                    "block": {
                      "block_types": {
                        "failure_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "failure_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "failure_response": {
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
                        "success_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "success_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "success_response": {
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
                        "timeout_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "timeout_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "timeout_response": {
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
            },
            "confirmation_conditional": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "conditional_branch": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "expression_string": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
                  },
                  "default_branch": {
                    "block": {
                      "block_types": {
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
            },
            "confirmation_next_step": {
              "block": {
                "attributes": {
                  "session_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dialog_action": {
                    "block": {
                      "attributes": {
                        "slot_to_elicit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suppress_next_message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "intent": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "slot": {
                          "block": {
                            "attributes": {
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "shape": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value": {
                                "block": {
                                  "attributes": {
                                    "interpreted_value": {
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
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "confirmation_response": {
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
            "declination_conditional": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "conditional_branch": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "expression_string": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
                  },
                  "default_branch": {
                    "block": {
                      "block_types": {
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
            },
            "declination_next_step": {
              "block": {
                "attributes": {
                  "session_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dialog_action": {
                    "block": {
                      "attributes": {
                        "slot_to_elicit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suppress_next_message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "intent": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "slot": {
                          "block": {
                            "attributes": {
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "shape": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value": {
                                "block": {
                                  "attributes": {
                                    "interpreted_value": {
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
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "declination_response": {
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
            "elicitation_code_hook": {
              "block": {
                "attributes": {
                  "enable_code_hook_invocation": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "invocation_label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "failure_conditional": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "conditional_branch": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "expression_string": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
                  },
                  "default_branch": {
                    "block": {
                      "block_types": {
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
            },
            "failure_next_step": {
              "block": {
                "attributes": {
                  "session_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dialog_action": {
                    "block": {
                      "attributes": {
                        "slot_to_elicit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suppress_next_message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "intent": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "slot": {
                          "block": {
                            "attributes": {
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "shape": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value": {
                                "block": {
                                  "attributes": {
                                    "interpreted_value": {
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
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "failure_response": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "dialog_code_hook": {
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
        "nesting_mode": "list"
      },
      "fulfillment_code_hook": {
        "block": {
          "attributes": {
            "active": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "fulfillment_updates_specification": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "timeout_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "start_response": {
                    "block": {
                      "attributes": {
                        "allow_interrupt": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delay_in_seconds": {
                          "description_kind": "plain",
                          "optional": true,
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
                  "update_response": {
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
            },
            "post_fulfillment_status_specification": {
              "block": {
                "block_types": {
                  "failure_conditional": {
                    "block": {
                      "attributes": {
                        "active": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "conditional_branch": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "condition": {
                                "block": {
                                  "attributes": {
                                    "expression_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                        },
                        "default_branch": {
                          "block": {
                            "block_types": {
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                  },
                  "failure_next_step": {
                    "block": {
                      "attributes": {
                        "session_attributes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "dialog_action": {
                          "block": {
                            "attributes": {
                              "slot_to_elicit": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suppress_next_message": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                        "intent": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "slot": {
                                "block": {
                                  "attributes": {
                                    "map_block_key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "shape": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "value": {
                                      "block": {
                                        "attributes": {
                                          "interpreted_value": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "failure_response": {
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
                  "success_conditional": {
                    "block": {
                      "attributes": {
                        "active": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "conditional_branch": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "condition": {
                                "block": {
                                  "attributes": {
                                    "expression_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                        },
                        "default_branch": {
                          "block": {
                            "block_types": {
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                  },
                  "success_next_step": {
                    "block": {
                      "attributes": {
                        "session_attributes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "dialog_action": {
                          "block": {
                            "attributes": {
                              "slot_to_elicit": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suppress_next_message": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                        "intent": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "slot": {
                                "block": {
                                  "attributes": {
                                    "map_block_key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "shape": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "value": {
                                      "block": {
                                        "attributes": {
                                          "interpreted_value": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "success_response": {
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
                  "timeout_conditional": {
                    "block": {
                      "attributes": {
                        "active": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "conditional_branch": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "condition": {
                                "block": {
                                  "attributes": {
                                    "expression_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                        },
                        "default_branch": {
                          "block": {
                            "block_types": {
                              "next_step": {
                                "block": {
                                  "attributes": {
                                    "session_attributes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "dialog_action": {
                                      "block": {
                                        "attributes": {
                                          "slot_to_elicit": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "suppress_next_message": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
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
                                    "intent": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "slot": {
                                            "block": {
                                              "attributes": {
                                                "map_block_key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "shape": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "interpreted_value": {
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
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "response": {
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
                  },
                  "timeout_next_step": {
                    "block": {
                      "attributes": {
                        "session_attributes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "dialog_action": {
                          "block": {
                            "attributes": {
                              "slot_to_elicit": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "suppress_next_message": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                        "intent": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "slot": {
                                "block": {
                                  "attributes": {
                                    "map_block_key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "shape": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "value": {
                                      "block": {
                                        "attributes": {
                                          "interpreted_value": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "timeout_response": {
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
      },
      "initial_response_setting": {
        "block": {
          "block_types": {
            "code_hook": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "enable_code_hook_invocation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "invocation_label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "post_code_hook_specification": {
                    "block": {
                      "block_types": {
                        "failure_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "failure_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "failure_response": {
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
                        "success_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "success_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "success_response": {
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
                        "timeout_conditional": {
                          "block": {
                            "attributes": {
                              "active": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "conditional_branch": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "condition": {
                                      "block": {
                                        "attributes": {
                                          "expression_string": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                              },
                              "default_branch": {
                                "block": {
                                  "block_types": {
                                    "next_step": {
                                      "block": {
                                        "attributes": {
                                          "session_attributes": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "dialog_action": {
                                            "block": {
                                              "attributes": {
                                                "slot_to_elicit": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "suppress_next_message": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
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
                                          "intent": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "slot": {
                                                  "block": {
                                                    "attributes": {
                                                      "map_block_key": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "shape": {
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "value": {
                                                        "block": {
                                                          "attributes": {
                                                            "interpreted_value": {
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
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    },
                                    "response": {
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
                        },
                        "timeout_next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "timeout_response": {
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
            },
            "conditional": {
              "block": {
                "attributes": {
                  "active": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "conditional_branch": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "expression_string": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
                  },
                  "default_branch": {
                    "block": {
                      "block_types": {
                        "next_step": {
                          "block": {
                            "attributes": {
                              "session_attributes": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "block_types": {
                              "dialog_action": {
                                "block": {
                                  "attributes": {
                                    "slot_to_elicit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suppress_next_message": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
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
                              "intent": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "slot": {
                                      "block": {
                                        "attributes": {
                                          "map_block_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "shape": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "value": {
                                            "block": {
                                              "attributes": {
                                                "interpreted_value": {
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
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "response": {
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
            },
            "initial_response": {
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
            "next_step": {
              "block": {
                "attributes": {
                  "session_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "dialog_action": {
                    "block": {
                      "attributes": {
                        "slot_to_elicit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "suppress_next_message": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "intent": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "slot": {
                          "block": {
                            "attributes": {
                              "map_block_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "shape": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "value": {
                                "block": {
                                  "attributes": {
                                    "interpreted_value": {
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
      "input_context": {
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
      "kendra_configuration": {
        "block": {
          "attributes": {
            "kendra_index": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "query_filter_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_filter_string_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "output_context": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "time_to_live_in_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "turns_to_live": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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
      "slot_priority": {
        "block": {
          "attributes": {
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "slot_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AwsLexv2ModelsIntentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLexv2ModelsIntent), &result)
	return &result
}
