package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftDataShareConsumerAssociation = `{
  "block": {
    "attributes": {
      "allow_writes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "associate_entire_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "consumer_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "consumer_region": {
        "description_kind": "plain",
        "optional": true,
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

func AwsRedshiftDataShareConsumerAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftDataShareConsumerAssociation), &result)
	return &result
}
