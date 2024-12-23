package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecretsmanagerSecretVersions = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "include_deprecated": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_time": "string",
              "last_accessed_date": "string",
              "version_id": "string",
              "version_stages": [
                "list",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecretsmanagerSecretVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecretsmanagerSecretVersions), &result)
	return &result
}
