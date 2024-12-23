package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmcontactsRotation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "contact_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
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
        "type": [
          "map",
          "string"
        ]
      },
      "time_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "recurrence": {
        "block": {
          "attributes": {
            "number_of_on_calls": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "recurrence_multiplier": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "daily_settings": {
              "block": {
                "attributes": {
                  "hour_of_day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minute_of_hour": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "monthly_settings": {
              "block": {
                "attributes": {
                  "day_of_month": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "hand_off_time": {
                    "block": {
                      "attributes": {
                        "hour_of_day": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "minute_of_hour": {
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
            "shift_coverages": {
              "block": {
                "attributes": {
                  "map_block_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "coverage_times": {
                    "block": {
                      "block_types": {
                        "end": {
                          "block": {
                            "attributes": {
                              "hour_of_day": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "minute_of_hour": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "start": {
                          "block": {
                            "attributes": {
                              "hour_of_day": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "minute_of_hour": {
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "weekly_settings": {
              "block": {
                "attributes": {
                  "day_of_week": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "hand_off_time": {
                    "block": {
                      "attributes": {
                        "hour_of_day": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "minute_of_hour": {
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

func AwsSsmcontactsRotationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmcontactsRotation), &result)
	return &result
}
