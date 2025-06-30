package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsTaskDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn_without_revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_definitions": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_fault_injection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ephemeral_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "size_in_gib": "number"
            }
          ]
        ]
      },
      "execution_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "family": {
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
      "ipc_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pid_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "placement_constraints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "expression": "string",
              "type": "string"
            }
          ]
        ]
      },
      "proxy_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_name": "string",
              "properties": [
                "map",
                "string"
              ],
              "type": "string"
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
      "requires_compatibilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "runtime_platform": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_architecture": "string",
              "operating_system_family": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_definition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "configure_at_launch": "bool",
              "docker_volume_configuration": [
                "list",
                [
                  "object",
                  {
                    "autoprovision": "bool",
                    "driver": "string",
                    "driver_opts": [
                      "map",
                      "string"
                    ],
                    "labels": [
                      "map",
                      "string"
                    ],
                    "scope": "string"
                  }
                ]
              ],
              "efs_volume_configuration": [
                "list",
                [
                  "object",
                  {
                    "authorization_config": [
                      "list",
                      [
                        "object",
                        {
                          "access_point_id": "string",
                          "iam": "string"
                        }
                      ]
                    ],
                    "file_system_id": "string",
                    "root_directory": "string",
                    "transit_encryption": "string",
                    "transit_encryption_port": "number"
                  }
                ]
              ],
              "fsx_windows_file_server_volume_configuration": [
                "list",
                [
                  "object",
                  {
                    "authorization_config": [
                      "list",
                      [
                        "object",
                        {
                          "credentials_parameter": "string",
                          "domain": "string"
                        }
                      ]
                    ],
                    "file_system_id": "string",
                    "root_directory": "string"
                  }
                ]
              ],
              "host_path": "string",
              "name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEcsTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsTaskDefinition), &result)
	return &result
}
