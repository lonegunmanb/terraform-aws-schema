package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsContainerDefinition = `{
  "block": {
    "attributes": {
      "container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "disable_networking": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "docker_labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "environment": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_digest": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "memory_reservation": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_definition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsContainerDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsContainerDefinition), &result)
	return &result
}
