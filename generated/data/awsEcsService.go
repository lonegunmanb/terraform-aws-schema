package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsService = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_rebalancing": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desired_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "advanced_configuration": [
                "list",
                [
                  "object",
                  {
                    "alternate_target_group_arn": "string",
                    "production_listener_rule": "string",
                    "role_arn": "string",
                    "test_listener_rule": "string"
                  }
                ]
              ],
              "container_name": "string",
              "container_port": "number",
              "elb_name": "string",
              "target_group_arn": "string"
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
      },
      "scheduling_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "task_definition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsService), &result)
	return &result
}
