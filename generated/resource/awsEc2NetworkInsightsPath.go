package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2NetworkInsightsPath = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_ip": {
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
      "filter_at_destination": {
        "block": {
          "attributes": {
            "destination_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination_port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
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
            "source_port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "filter_at_source": {
        "block": {
          "attributes": {
            "destination_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination_port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
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
            "source_port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
  "version": 0
}`

func AwsEc2NetworkInsightsPathSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2NetworkInsightsPath), &result)
	return &result
}
