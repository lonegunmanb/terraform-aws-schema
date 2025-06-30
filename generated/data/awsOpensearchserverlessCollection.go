package data

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
      "created_date": {
        "computed": true,
        "description": "Date the Collection was created.",
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
        "computed": true,
        "description": "Description of the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_code": {
        "computed": true,
        "description": "A failure code associated with the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_message": {
        "computed": true,
        "description": "A failure reason associated with the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "ID of the collection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the Amazon Web Services KMS key used to encrypt the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "Date the Collection was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the collection.",
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
      "standby_replicas": {
        "computed": true,
        "description": "Indicates whether standby replicas should be used for a collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "Type of collection.",
        "description_kind": "plain",
        "type": "string"
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
