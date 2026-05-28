package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3ControlMultiRegionAccessPointRoutes = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mrap": {
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
      }
    },
    "block_types": {
      "route": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "traffic_dial_percentage": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3ControlMultiRegionAccessPointRoutesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3ControlMultiRegionAccessPointRoutes), &result)
	return &result
}
