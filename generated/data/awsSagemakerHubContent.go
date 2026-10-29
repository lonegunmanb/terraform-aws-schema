package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerHubContent = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "document_schema_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_dependency": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dependency_copy_path": "string",
              "dependency_origin_path": "string"
            }
          ]
        ]
      },
      "hub_content_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_document": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_markdown": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hub_content_search_keywords": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "hub_content_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hub_content_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hub_content_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hub_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reference_min_version": {
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
      "sagemaker_public_hub_content_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "versioned_hub_content_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerHubContentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerHubContent), &result)
	return &result
}
