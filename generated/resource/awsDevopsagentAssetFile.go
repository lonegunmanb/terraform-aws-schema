package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDevopsagentAssetFile = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "description": "Identifier of the Agent Space containing the asset. Changing this value requires replacement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_id": {
        "description": "Identifier of the asset containing the file. Changing this value requires replacement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Creation timestamp in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "description": "JSON document describing the file. Updates replace supplied top-level fields and retain omitted fields in AWS; removing a field from this document does not delete it from AWS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata_actual": {
        "computed": true,
        "description": "Full JSON metadata returned by AWS, including retained fields and service-provided defaults.",
        "description_kind": "plain",
        "type": "string"
      },
      "path": {
        "description": "Path of the file within the asset. Changing this value requires replacement.",
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
      "updated_at": {
        "computed": true,
        "description": "Last update timestamp in RFC 3339 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Asset version the file belongs to.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "content": {
        "block": {
          "attributes": {
            "bytes": {
              "description": "Base64-encoded binary file content.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "text": {
              "description": "Text file content.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "File content. Exactly one block is required.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDevopsagentAssetFileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDevopsagentAssetFile), &result)
	return &result
}
