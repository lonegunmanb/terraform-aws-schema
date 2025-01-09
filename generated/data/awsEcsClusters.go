package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsClusters = `{
  "block": {
    "attributes": {
      "cluster_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsClusters), &result)
	return &result
}
