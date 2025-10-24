package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOrganizationsOrganization = `{
  "block": {
    "attributes": {
      "accounts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "email": "string",
              "id": "string",
              "joined_method": "string",
              "joined_timestamp": "string",
              "name": "string",
              "state": "string",
              "status": "string"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_service_access_principals": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "enabled_policy_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "feature_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_account_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "non_master_accounts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "email": "string",
              "id": "string",
              "joined_method": "string",
              "joined_timestamp": "string",
              "name": "string",
              "state": "string",
              "status": "string"
            }
          ]
        ]
      },
      "roots": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "id": "string",
              "name": "string",
              "policy_types": [
                "list",
                [
                  "object",
                  {
                    "status": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOrganizationsOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOrganizationsOrganization), &result)
	return &result
}
