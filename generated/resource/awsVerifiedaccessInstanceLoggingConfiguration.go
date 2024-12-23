package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedaccessInstanceLoggingConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "verifiedaccess_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "access_logs": {
        "block": {
          "attributes": {
            "include_trust_context": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "log_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logs": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "log_group": {
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
            "kinesis_data_firehose": {
              "block": {
                "attributes": {
                  "delivery_stream": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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
            },
            "s3": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "prefix": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedaccessInstanceLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedaccessInstanceLoggingConfiguration), &result)
	return &result
}
