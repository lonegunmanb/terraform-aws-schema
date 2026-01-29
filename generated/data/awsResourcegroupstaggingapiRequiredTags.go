package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsResourcegroupstaggingapiRequiredTags = `{
  "block": {
    "attributes": {
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "required_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloud_formation_resource_types": [
                "list",
                "string"
              ],
              "reporting_tag_keys": [
                "list",
                "string"
              ],
              "resource_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsResourcegroupstaggingapiRequiredTagsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsResourcegroupstaggingapiRequiredTags), &result)
	return &result
}
