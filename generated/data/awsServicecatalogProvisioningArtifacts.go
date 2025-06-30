package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogProvisioningArtifacts = `{
  "block": {
    "attributes": {
      "accept_language": {
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
      "product_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioning_artifact_details": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active": "bool",
              "created_time": "string",
              "description": "string",
              "guidance": "string",
              "id": "string",
              "name": "string",
              "type": "string"
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
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogProvisioningArtifactsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogProvisioningArtifacts), &result)
	return &result
}
