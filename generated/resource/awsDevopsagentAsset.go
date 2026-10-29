package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDevopsagentAsset = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "description": "Identifier of the Agent Space containing the asset. Changing this value requires replacement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_id": {
        "computed": true,
        "description": "Identifier of the asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_type": {
        "description": "Type of asset. Changing this value requires replacement. See https://docs.aws.amazon.com/devopsagent/latest/userguide/about-aws-devops-agent-managing-assets.html for supported types and their metadata requirements.",
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
        "description": "JSON document describing the asset. Required metadata fields depend on asset_type. Updates replace supplied top-level fields and retain omitted fields in AWS; removing a field from this document does not delete it from AWS. Changes to metadata alone do not re-upload content.",
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
        "description": "Current asset version.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "content": {
        "block": {
          "block_types": {
            "file": {
              "block": {
                "attributes": {
                  "metadata": {
                    "description": "JSON metadata for the file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description": "Path of the file within the asset.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "body": {
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
                      "description": "File content.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Single file to add or replace.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "source_url": {
              "block": {
                "attributes": {
                  "url": {
                    "description": "Source URL from which to import asset content.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Source URL to import.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "zip": {
              "block": {
                "attributes": {
                  "zip_file": {
                    "description": "Base64-encoded ZIP archive. Use filebase64() to load a local ZIP file.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "ZIP archive containing asset files.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Content to upload. Exactly one block is required.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDevopsagentAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDevopsagentAsset), &result)
	return &result
}
