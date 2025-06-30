package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrSupportedInstanceTypes = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_label": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "supported_instance_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "architecture": "string",
              "ebs_optimized_available": "bool",
              "ebs_optimized_by_default": "bool",
              "ebs_storage_only": "bool",
              "instance_family_id": "string",
              "is_64_bits_only": "bool",
              "memory_gb": "number",
              "number_of_disks": "number",
              "storage_gb": "number",
              "type": "string",
              "vcpu": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEmrSupportedInstanceTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrSupportedInstanceTypes), &result)
	return &result
}
