package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubConnectorV2 = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connector_status": "string",
              "last_checked_at": "string",
              "message": "string"
            }
          ]
        ]
      },
      "kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "connector_provider": {
        "block": {
          "block_types": {
            "jira_cloud": {
              "block": {
                "attributes": {
                  "auth_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auth_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cloud_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "project_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "service_now": {
              "block": {
                "attributes": {
                  "auth_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubConnectorV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubConnectorV2), &result)
	return &result
}
