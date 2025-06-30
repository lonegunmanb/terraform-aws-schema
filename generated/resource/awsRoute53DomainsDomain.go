package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53DomainsDomain = `{
  "block": {
    "attributes": {
      "abuse_contact_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "abuse_contact_phone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "admin_privacy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "billing_contact": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "address_line_1": "string",
              "address_line_2": "string",
              "city": "string",
              "contact_type": "string",
              "country_code": "string",
              "email": "string",
              "extra_param": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "value": "string"
                  }
                ]
              ],
              "fax": "string",
              "first_name": "string",
              "last_name": "string",
              "organization_name": "string",
              "phone_number": "string",
              "state": "string",
              "zip_code": "string"
            }
          ]
        ]
      },
      "billing_privacy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration_in_years": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "expiration_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name_server": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "glue_ips": [
                "set",
                "string"
              ],
              "name": "string"
            }
          ]
        ]
      },
      "registrant_privacy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "registrar_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "registrar_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
      },
      "tech_privacy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transfer_lock": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "whois_server": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "admin_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fax": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "extra_param": {
              "block": {
                "attributes": {
                  "name": {
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
      "registrant_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fax": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "extra_param": {
              "block": {
                "attributes": {
                  "name": {
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
      "tech_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fax": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "extra_param": {
              "block": {
                "attributes": {
                  "name": {
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

func AwsRoute53DomainsDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53DomainsDomain), &result)
	return &result
}
