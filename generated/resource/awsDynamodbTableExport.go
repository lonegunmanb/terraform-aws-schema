package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDynamodbTableExport = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "billed_size_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "export_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "export_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "export_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "export_type": {
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
      "item_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "manifest_files_s3_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket_owner": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_sse_algorithm": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_sse_kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "incremental_export_specification": {
        "block": {
          "attributes": {
            "export_from_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_view_type": {
              "computed": true,
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsDynamodbTableExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDynamodbTableExport), &result)
	return &result
}
