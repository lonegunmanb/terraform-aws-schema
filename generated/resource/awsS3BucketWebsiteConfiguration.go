package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketWebsiteConfiguration = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expected_bucket_owner": {
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
      "routing_rules": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "website_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "website_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "error_document": {
        "block": {
          "attributes": {
            "key": {
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
      "index_document": {
        "block": {
          "attributes": {
            "suffix": {
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
      "redirect_all_requests_to": {
        "block": {
          "attributes": {
            "host_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
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
      "routing_rule": {
        "block": {
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "http_error_code_returned_equals": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_prefix_equals": {
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
            "redirect": {
              "block": {
                "attributes": {
                  "host_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_redirect_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "replace_key_prefix_with": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "replace_key_with": {
                    "description_kind": "plain",
                    "optional": true,
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketWebsiteConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketWebsiteConfiguration), &result)
	return &result
}
