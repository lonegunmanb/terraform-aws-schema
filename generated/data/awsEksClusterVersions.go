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
