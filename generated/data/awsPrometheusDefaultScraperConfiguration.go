package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPrometheusDefaultScraperConfiguration = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsPrometheusDefaultScraperConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPrometheusDefaultScraperConfiguration), &result)
	return &result
}
