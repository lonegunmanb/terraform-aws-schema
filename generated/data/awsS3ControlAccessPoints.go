package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3ControlAccessPoints = `{
  "block": {
    "attributes": {
      "access_points": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_point_arn": "string",
              "alias": "string",
              "bucket": "string",
              "bucket_account_id": "string",
              "data_source_id": "string",
              "data_source_type": "string",
              "name": "string",
              "network_origin": "string",
              "vpc_configuration": [
                "list",
                [
                  "object",
                  {
                    "vpc_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_type": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3ControlAccessPointsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3ControlAccessPoints), &result)
	return &result
}
