package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessCollection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collection_endpoint": {
        "computed": true,
        "description": "Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_endpoint": {
        "computed": true,
        "description": "Collection-specific endpoint used to access OpenSearch Dashboards.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the collection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the Amazon Web Services KMS key used to encrypt the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the collection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "standby_replicas": {
        "computed": true,
        "description": "Indicates whether standby replicas should be used for a collection. One of ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. Defaults to ` + "`" + `ENABLED` + "`" + `.",
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
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "Type of collection. One of ` + "`" + `SEARCH` + "`" + `, ` + "`" + `TIMESERIES` + "`" + `, or ` + "`" + `VECTORSEARCH` + "`" + `. Defaults to ` + "`" + `TIMESERIES` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AwsOpensearchserverlessCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessCollection), &result)
	return &result
}
