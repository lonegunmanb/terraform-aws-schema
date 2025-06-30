package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "broker_node_group_info": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "az_distribution": "string",
              "client_subnets": [
                "set",
                "string"
              ],
              "connectivity_info": [
                "list",
                [
                  "object",
                  {
                    "public_access": [
                      "list",
                      [
                        "object",
                        {
                          "type": "string"
                        }
                      ]
                    ],
                    "vpc_connectivity": [
                      "list",
                      [
                        "object",
                        {
                          "client_authentication": [
                            "list",
                            [
                              "object",
                              {
                                "sasl": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "iam": "bool",
                                      "scram": "bool"
                                    }
                                  ]
                                ],
                                "tls": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "instance_type": "string",
              "security_groups": [
                "set",
                "string"
              ],
              "storage_info": [
                "list",
                [
                  "object",
                  {
                    "ebs_storage_info": [
                      "list",
                      [
                        "object",
                        {
                          "provisioned_throughput": [
                            "list",
                            [
                              "object",
                              {
                                "enabled": "bool",
                                "volume_throughput": "number"
                              }
                            ]
                          ],
                          "volume_size": "number"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_uuid": {
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
      "kafka_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_broker_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      },
      "zookeeper_connect_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zookeeper_connect_string_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMskClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskCluster), &result)
	return &result
}
