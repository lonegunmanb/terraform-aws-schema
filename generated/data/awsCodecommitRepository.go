package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodecommitRepository = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "clone_url_http": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "clone_url_ssh": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "repository_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "repository_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodecommitRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodecommitRepository), &result)
	return &result
}
