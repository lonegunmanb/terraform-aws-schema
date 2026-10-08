package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdamicrovmsImageVersion = `{
  "block": {
    "attributes": {
      "additional_os_capabilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "base_image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "base_image_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "code_artifact": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "uri": "string"
            }
          ]
        ]
      },
      "cpu_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "architecture": "string"
            }
          ]
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "egress_network_connectors": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "environment_variables": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "hooks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "microvm_hooks": [
                "list",
                [
                  "object",
                  {
                    "resume": "string",
                    "resume_timeout_in_seconds": "number",
                    "run": "string",
                    "run_timeout_in_seconds": "number",
                    "suspend": "string",
                    "suspend_timeout_in_seconds": "number",
                    "terminate": "string",
                    "terminate_timeout_in_seconds": "number"
                  }
                ]
              ],
              "microvm_image_hooks": [
                "list",
                [
                  "object",
                  {
                    "ready": "string",
                    "ready_timeout_in_seconds": "number",
                    "validate": "string",
                    "validate_timeout_in_seconds": "number"
                  }
                ]
              ],
              "port": "number"
            }
          ]
        ]
      },
      "image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logging": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cloudwatch": [
                "list",
                [
                  "object",
                  {
                    "log_group": "string",
                    "log_stream": "string"
                  }
                ]
              ],
              "disabled": [
                "list",
                [
                  "object",
                  {}
                ]
              ]
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
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "minimum_memory_in_mib": "number"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdamicrovmsImageVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdamicrovmsImageVersion), &result)
	return &result
}
