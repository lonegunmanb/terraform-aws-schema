package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEksClusterVersions = `{
  "block": {
    "attributes": {
      "cluster_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_type": "string",
              "cluster_version": "string",
              "control_plane_component_config": [
                "list",
                [
                  "object",
                  {
                    "kube_api_server_config": [
                      "list",
                      [
                        "object",
                        {
                          "event_ttl": [
                            "list",
                            [
                              "object",
                              {
                                "constraints": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "string",
                                      "min": "string"
                                    }
                                  ]
                                ],
                                "default_value": "string"
                              }
                            ]
                          ],
                          "service_node_port_range": [
                            "list",
                            [
                              "object",
                              {
                                "constraints": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max_port": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max": "number",
                                            "min": "number"
                                          }
                                        ]
                                      ],
                                      "min_port": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max": "number",
                                            "min": "number"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "default_value": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max_port": "number",
                                      "min_port": "number"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "kube_controller_manager_config": [
                      "list",
                      [
                        "object",
                        {
                          "horizontal_pod_autoscaler_controller_config": [
                            "list",
                            [
                              "object",
                              {
                                "horizontal_pod_autoscaler_sync_period": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "constraints": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max": "string",
                                            "min": "string"
                                          }
                                        ]
                                      ],
                                      "default_value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "kube_scheduler_config": [
                      "list",
                      [
                        "object",
                        {
                          "node_resources_fit": [
                            "list",
                            [
                              "object",
                              {
                                "scoring_strategy": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "constraints": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "resources": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "name": [
                                                    "list",
                                                    [
                                                      "object",
                                                      {
                                                        "allowed_values": [
                                                          "list",
                                                          "string"
                                                        ]
                                                      }
                                                    ]
                                                  ],
                                                  "weight": [
                                                    "list",
                                                    [
                                                      "object",
                                                      {
                                                        "max": "number",
                                                        "min": "number"
                                                      }
                                                    ]
                                                  ]
                                                }
                                              ]
                                            ],
                                            "scoring_strategy": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "allowed_values": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ]
                                          }
                                        ]
                                      ],
                                      "default_value": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "resources": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "name": "string",
                                                  "weight": "number"
                                                }
                                              ]
                                            ],
                                            "type": "string"
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
              "control_plane_scaling_tiers": [
                "list",
                [
                  "object",
                  {
                    "api_request_concurrency": "number",
                    "cluster_database_size_gb": "number",
                    "control_plane_component_config_overrides": [
                      "list",
                      [
                        "object",
                        {
                          "kube_api_server_config": [
                            "list",
                            [
                              "object",
                              {
                                "event_ttl": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "constraints": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max": "string",
                                            "min": "string"
                                          }
                                        ]
                                      ],
                                      "default_value": "string"
                                    }
                                  ]
                                ],
                                "service_node_port_range": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "constraints": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max_port": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "max": "number",
                                                  "min": "number"
                                                }
                                              ]
                                            ],
                                            "min_port": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "max": "number",
                                                  "min": "number"
                                                }
                                              ]
                                            ]
                                          }
                                        ]
                                      ],
                                      "default_value": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "max_port": "number",
                                            "min_port": "number"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "kube_controller_manager_config": [
                            "list",
                            [
                              "object",
                              {
                                "horizontal_pod_autoscaler_controller_config": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "horizontal_pod_autoscaler_sync_period": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "constraints": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "max": "string",
                                                  "min": "string"
                                                }
                                              ]
                                            ],
                                            "default_value": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "kube_scheduler_config": [
                            "list",
                            [
                              "object",
                              {
                                "node_resources_fit": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "scoring_strategy": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "constraints": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "resources": [
                                                    "list",
                                                    [
                                                      "object",
                                                      {
                                                        "name": [
                                                          "list",
                                                          [
                                                            "object",
                                                            {
                                                              "allowed_values": [
                                                                "list",
                                                                "string"
                                                              ]
                                                            }
                                                          ]
                                                        ],
                                                        "weight": [
                                                          "list",
                                                          [
                                                            "object",
                                                            {
                                                              "max": "number",
                                                              "min": "number"
                                                            }
                                                          ]
                                                        ]
                                                      }
                                                    ]
                                                  ],
                                                  "scoring_strategy": [
                                                    "list",
                                                    [
                                                      "object",
                                                      {
                                                        "allowed_values": [
                                                          "list",
                                                          "string"
                                                        ]
                                                      }
                                                    ]
                                                  ]
                                                }
                                              ]
                                            ],
                                            "default_value": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "resources": [
                                                    "list",
                                                    [
                                                      "object",
                                                      {
                                                        "name": "string",
                                                        "weight": "number"
                                                      }
                                                    ]
                                                  ],
                                                  "type": "string"
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
                    "pod_scheduling_rate_per_second": "number",
                    "tier_name": "string"
                  }
                ]
              ],
              "default_platform_version": "string",
              "default_version": "bool",
              "end_of_extended_support_date": "string",
              "end_of_standard_support_date": "string",
              "kubernetes_patch_version": "string",
              "release_date": "string",
              "version_status": "string"
            }
          ]
        ]
      },
      "cluster_versions_only": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "default_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_all": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEksClusterVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEksClusterVersions), &result)
	return &result
}
