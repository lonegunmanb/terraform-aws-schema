package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEksCluster = `{
  "block": {
    "attributes": {
      "access_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authentication_mode": "string",
              "bootstrap_cluster_creator_admin_permissions": "bool"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "node_pools": [
                "set",
                "string"
              ],
              "node_role_arn": "string"
            }
          ]
        ]
      },
      "control_plane_scaling_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "tier": "string"
            }
          ]
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_cluster_log_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "endpoint": {
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "oidc": [
                "list",
                [
                  "object",
                  {
                    "issuer": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "kubernetes_network_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "elastic_load_balancing": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "ip_family": "string",
              "service_ipv4_cidr": "string",
              "service_ipv6_cidr": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outpost_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "control_plane_instance_type": "string",
              "control_plane_placement": [
                "list",
                [
                  "object",
                  {
                    "group_name": "string"
                  }
                ]
              ],
              "outpost_arns": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "platform_version": {
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
      "remote_network_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "remote_node_networks": [
                "list",
                [
                  "object",
                  {
                    "cidrs": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "remote_pod_networks": [
                "list",
                [
                  "object",
                  {
                    "cidrs": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "block_storage": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ]
            }
          ]
        ]
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
      "upgrade_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "support_type": "string"
            }
          ]
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_security_group_id": "string",
              "endpoint_private_access": "bool",
              "endpoint_public_access": "bool",
              "public_access_cidrs": [
                "set",
                "string"
              ],
              "security_group_ids": [
                "set",
                "string"
              ],
              "subnet_ids": [
                "set",
                "string"
              ],
              "vpc_id": "string"
            }
          ]
        ]
      },
      "zonal_shift_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEksClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEksCluster), &result)
	return &result
}
