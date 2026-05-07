package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "athena_properties": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "authentication_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authentication_type": "string",
              "basic_authentication_credentials": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "username": "string"
                  }
                ]
              ],
              "custom_authentication_credentials": [
                "map",
                "string"
              ],
              "kms_key_arn": "string",
              "oauth2_properties": [
                "list",
                [
                  "object",
                  {
                    "authorization_code_properties": [
                      "list",
                      [
                        "object",
                        {
                          "authorization_code": "string",
                          "redirect_uri": "string"
                        }
                      ]
                    ],
                    "oauth2_client_application": [
                      "list",
                      [
                        "object",
                        {
                          "aws_managed_client_application_reference": "string",
                          "user_managed_client_application_client_id": "string"
                        }
                      ]
                    ],
                    "oauth2_credentials": [
                      "list",
                      [
                        "object",
                        {
                          "access_token": "string",
                          "jwt_token": "string",
                          "refresh_token": "string",
                          "user_managed_client_application_client_secret": "string"
                        }
                      ]
                    ],
                    "oauth2_grant_type": "string",
                    "token_url": "string",
                    "token_url_parameters_map": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "secret_arn": "string"
            }
          ]
        ]
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_properties": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "connection_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "match_criteria": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "physical_connection_requirements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zone": "string",
              "security_group_id_list": [
                "set",
                "string"
              ],
              "subnet_id": "string"
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueConnection), &result)
	return &result
}
