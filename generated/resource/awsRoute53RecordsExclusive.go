package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53RecordsExclusive = `{
  "block": {
    "attributes": {
      "zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "resource_record_set": {
        "block": {
          "attributes": {
            "failover": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "multi_value_answer": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "set_identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "traffic_policy_instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "alias_target": {
              "block": {
                "attributes": {
                  "dns_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "hosted_zone_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cidr_routing_config": {
              "block": {
                "attributes": {
                  "collection_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "location_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "geolocation": {
              "block": {
                "attributes": {
                  "continent_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "country_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subdivision_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "geoproximity_location": {
              "block": {
                "attributes": {
                  "aws_region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "local_zone_group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "coordinates": {
                    "block": {
                      "attributes": {
                        "latitude": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "longitude": {
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
            "resource_records": {
              "block": {
                "attributes": {
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
        "nesting_mode": "set"
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

func AwsRoute53RecordsExclusiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53RecordsExclusive), &result)
	return &result
}
