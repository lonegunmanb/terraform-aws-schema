package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmAssociation = `{
  "block": {
    "attributes": {
      "apply_only_at_cron_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "association_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automation_target_parameter_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compliance_severity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "document_version": {
        "computed": true,
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
      "max_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_errors": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sync_compliance": {
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
      },
      "wait_for_success_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "output_location": {
        "block": {
          "attributes": {
            "s3_bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_region": {
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
      "targets": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsSsmAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmAssociation), &result)
	return &result
}
