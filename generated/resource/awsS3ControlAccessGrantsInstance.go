package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3ControlAccessGrantsInstance = `{
  "block": {
    "attributes": {
      "access_grants_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_grants_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3ControlAccessGrantsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3ControlAccessGrantsInstance), &result)
	return &result
}
