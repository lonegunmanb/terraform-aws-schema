package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftDataShareAuthorization = `{
  "block": {
    "attributes": {
      "allow_writes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "consumer_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_share_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "managed_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "producer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftDataShareAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftDataShareAuthorization), &result)
	return &result
}