package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAgentregistryRegistry = `{
  "block": {
    "attributes": {
      "approval_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_approval_rules": [
                "set",
                "string"
              ]
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
      "discovery_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authorizer_configuration": [
                "list",
                [
                  "object",
                  {
                    "custom_jwt_authorizer": [
                      "list",
                      [
                        "object",
                        {
                          "allowed_audience": [
                            "list",
                            "string"
                          ],
                          "allowed_clients": [
                            "list",
                            "string"
                          ],
                          "allowed_scopes": [
                            "list",
                            "string"
                          ],
                          "custom_claim": [
                            "set",
                            [
                              "object",
                              {
                                "authorizing_claim_match_value": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claim_match_operator": "string",
                                      "claim_match_value": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "match_value_string": "string",
                                            "match_value_string_list": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "inbound_token_claim_name": "string",
                                "inbound_token_claim_value_type": "string"
                              }
                            ]
                          ],
                          "discovery_url": "string",
                          "private_endpoint": [
                            "list",
                            [
                              "object",
                              {
                                "managed_vpc_resource": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "endpoint_ip_address_type": "string",
                                      "routing_domain": "string",
                                      "security_group_ids": [
                                        "set",
                                        "string"
                                      ],
                                      "subnet_ids": [
                                        "set",
                                        "string"
                                      ],
                                      "tags": [
                                        "map",
                                        "string"
                                      ],
                                      "vpc_identifier": "string"
                                    }
                                  ]
                                ],
                                "self_managed_lattice_resource": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "resource_configuration_identifier": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "private_endpoint_override": [
                            "list",
                            [
                              "object",
                              {
                                "domain": "string",
                                "private_endpoint": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "managed_vpc_resource": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "endpoint_ip_address_type": "string",
                                            "routing_domain": "string",
                                            "security_group_ids": [
                                              "set",
                                              "string"
                                            ],
                                            "subnet_ids": [
                                              "set",
                                              "string"
                                            ],
                                            "tags": [
                                              "map",
                                              "string"
                                            ],
                                            "vpc_identifier": "string"
                                          }
                                        ]
                                      ],
                                      "self_managed_lattice_resource": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "resource_configuration_identifier": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "authorizer_type": "string"
            }
          ]
        ]
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_arn": "string"
            }
          ]
        ]
      },
      "name": {
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
      "registry_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "description_kind": "plain",
        "required": true,
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

func AwsAgentregistryRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAgentregistryRegistry), &result)
	return &result
}
