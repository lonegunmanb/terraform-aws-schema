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
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
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
