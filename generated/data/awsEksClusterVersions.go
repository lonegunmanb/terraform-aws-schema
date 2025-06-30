package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEksClusterVersions = `{
  "block": {
    "attributes": {
      "cluster_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_type": "string",
              "cluster_version": "string",
              "default_platform_version": "string",
              "default_version": "bool",
              "end_of_extended_support_date": "string",
              "end_of_standard_support_date": "string",
              "kubernetes_patch_version": "string",
              "release_date": "string",
              "version_status": "string"
            }
          ]
        ]
      },
      "cluster_versions_only": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "default_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_all": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEksClusterVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEksClusterVersions), &result)
	return &result
}
