package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerModelCardExportJob = `{
  "block": {
    "attributes": {
      "export_artifacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "s3_export_artifacts": "string"
            }
          ]
        ]
      },
      "model_card_export_job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "model_card_export_job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_card_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_card_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "output_config": {
        "block": {
          "attributes": {
            "s3_output_path": {
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

func AwsSagemakerModelCardExportJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerModelCardExportJob), &result)
	return &result
}
