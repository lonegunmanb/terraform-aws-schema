package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMqBrokerEngineTypes = `{
  "block": {
    "attributes": {
      "broker_engine_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "engine_type": "string",
              "engine_versions": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "engine_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
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

func AwsMqBrokerEngineTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMqBrokerEngineTypes), &result)
	return &result
}
