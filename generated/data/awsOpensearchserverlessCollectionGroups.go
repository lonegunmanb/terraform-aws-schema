package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessCollectionGroups = `{
  "block": {
    "attributes": {
      "collection_group_summaries": {
        "computed": true,
        "description": "List of collection group summaries.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "capacity_limits": [
                "list",
                [
                  "object",
                  {
                    "max_indexing_capacity_in_ocu": "number",
                    "max_search_capacity_in_ocu": "number",
                    "min_indexing_capacity_in_ocu": "number",
                    "min_search_capacity_in_ocu": "number"
                  }
                ]
              ],
              "created_date": "string",
              "id": "string",
              "name": "string",
              "number_of_collections": "number",
              "standby_replicas": "string"
            }
          ]
        ]
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

func AwsOpensearchserverlessCollectionGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessCollectionGroups), &result)
	return &result
}
