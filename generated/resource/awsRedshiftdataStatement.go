package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftdataStatement = `{
  "block": {
    "attributes": {
      "cluster_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_user": {
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
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statement_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "with_event": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "workgroup_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "parameters": {
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
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AwsRedshiftdataStatementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftdataStatement), &result)
	return &result
}
