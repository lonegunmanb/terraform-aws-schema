package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBatchJobDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_orchestration_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "eks_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pod_properties": [
                "list",
                [
                  "object",
                  {
                    "containers": [
                      "list",
                      [
                        "object",
                        {
                          "args": [
                            "list",
                            "string"
                          ],
                          "command": [
                            "list",
                            "string"
                          ],
                          "env": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "image": "string",
                          "image_pull_policy": "string",
                          "name": "string",
                          "resources": [
                            "list",
                            [
                              "object",
                              {
                                "limits": [
                                  "map",
                                  "string"
                                ],
                                "requests": [
                                  "map",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "security_context": [
                            "list",
                            [
                              "object",
                              {
                                "privileged": "bool",
                                "read_only_root_file_system": "bool",
                                "run_as_group": "number",
                                "run_as_non_root": "bool",
                                "run_as_user": "number"
                              }
                            ]
                          ],
                          "volume_mounts": [
                            "list",
                            [
                              "object",
                              {
                                "mount_path": "string",
                                "name": "string",
                                "read_only": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "dns_policy": "string",
                    "host_network": "bool",
                    "metadata": [
                      "list",
                      [
                        "object",
                        {
                          "labels": [
                            "map",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "service_account_name": "bool",
                    "volumes": [
                      "list",
                      [
                        "object",
                        {
                          "empty_dir": [
                            "list",
                            [
                              "object",
                              {
                                "medium": "string",
                                "size_limit": "string"
                              }
                            ]
                          ],
                          "host_path": [
                            "list",
                            [
                              "object",
                              {
                                "path": "string"
                              }
                            ]
                          ],
                          "name": "string",
                          "secret": [
                            "list",
                            [
                              "object",
                              {
                                "optional": "bool",
                                "secret_name": "string"
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
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "main_node": "number",
              "node_range_properties": [
                "list",
                [
                  "object",
                  {
                    "container": [
                      "list",
                      [
                        "object",
                        {
                          "command": [
                            "list",
                            "string"
                          ],
                          "environment": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "ephemeral_storage": [
                            "list",
                            [
                              "object",
                              {
                                "size_in_gib": "number"
                              }
                            ]
                          ],
                          "execution_role_arn": "string",
                          "fargate_platform_configuration": [
                            "list",
                            [
                              "object",
                              {
                                "platform_version": "string"
                              }
                            ]
                          ],
                          "image": "string",
                          "instance_type": "string",
                          "job_role_arn": "string",
                          "linux_parameters": [
                            "list",
                            [
                              "object",
                              {
                                "devices": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "container_path": "string",
                                      "host_path": "string",
                                      "permissions": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  ]
                                ],
                                "init_process_enabled": "bool",
                                "max_swap": "number",
                                "shared_memory_size": "number",
                                "swappiness": "number",
                                "tmpfs": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "container_path": "string",
                                      "mount_options": [
                                        "list",
                                        "string"
                                      ],
                                      "size": "number"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "log_configuration": [
                            "list",
                            [
                              "object",
                              {
                                "log_driver": "string",
                                "options": [
                                  "map",
                                  "string"
                                ],
                                "secret_options": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value_from": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "mount_points": [
                            "list",
                            [
                              "object",
                              {
                                "container_path": "string",
                                "read_only": "bool",
                                "source_volume": "string"
                              }
                            ]
                          ],
                          "network_configuration": [
                            "list",
                            [
                              "object",
                              {
                                "assign_public_ip": "bool"
                              }
                            ]
                          ],
                          "privileged": "bool",
                          "readonly_root_filesystem": "bool",
                          "resource_requirements": [
                            "list",
                            [
                              "object",
                              {
                                "type": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "runtime_platform": [
                            "list",
                            [
                              "object",
                              {
                                "cpu_architecture": "string",
                                "operating_system_family": "string"
                              }
                            ]
                          ],
                          "secrets": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value_from": "string"
                              }
                            ]
                          ],
                          "ulimits": [
                            "list",
                            [
                              "object",
                              {
                                "hard_limit": "number",
                                "name": "string",
                                "soft_limit": "number"
                              }
                            ]
                          ],
                          "user": "string",
                          "volumes": [
                            "list",
                            [
                              "object",
                              {
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
                                "host": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "source_path": "string"
                                    }
                                  ]
                                ],
                                "name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "target_nodes": "string"
                  }
                ]
              ],
              "num_nodes": "number"
            }
          ]
        ]
      },
      "retry_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attempts": "number",
              "evaluate_on_exit": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "on_exit_code": "string",
                    "on_reason": "string",
                    "on_status_reason": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "revision": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scheduling_priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
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
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attempt_duration_seconds": "number"
            }
          ]
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBatchJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBatchJobDefinition), &result)
	return &result
}
