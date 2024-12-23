package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCustomerprofilesProfile = `{
  "block": {
    "attributes": {
      "account_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "additional_information": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "birth_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "business_email_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "business_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "business_phone_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "first_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gender_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "home_phone_number": {
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
      "last_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "middle_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mobile_phone_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "party_type_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "personal_email_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "phone_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "address": {
        "block": {
          "attributes": {
            "address_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_3": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_4": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "county": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "province": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
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
      "billing_address": {
        "block": {
          "attributes": {
            "address_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_3": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_4": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "county": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "province": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
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
      "mailing_address": {
        "block": {
          "attributes": {
            "address_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_3": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_4": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "county": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "province": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
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
      "shipping_address": {
        "block": {
          "attributes": {
            "address_1": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_2": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_3": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_4": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "county": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "province": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
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
  "version": 0
}`

func AwsCustomerprofilesProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCustomerprofilesProfile), &result)
	return &result
}
