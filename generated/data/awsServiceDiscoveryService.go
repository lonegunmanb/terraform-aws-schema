package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServiceDiscoveryService = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dns_records": [
                "list",
                [
                  "object",
                  {
                    "ttl": "number",
                    "type": "string"
                  }
                ]
              ],
              "namespace_id": "string",
              "routing_policy": "string"
            }
          ]
        ]
      },
      "health_check_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "failure_threshold": "number",
              "resource_path": "string",
              "type": "string"
            }
          ]
        ]
      },
      "health_check_custom_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "failure_threshold": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_id": {
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
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServiceDiscoveryServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServiceDiscoveryService), &result)
	return &result
}
