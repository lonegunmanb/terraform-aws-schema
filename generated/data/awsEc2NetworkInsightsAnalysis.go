package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2NetworkInsightsAnalysis = `{
  "block": {
    "attributes": {
      "alternate_path_hints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "component_arn": "string",
              "component_id": "string"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "explanations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "acl_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "egress": "bool",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "rule_action": "string",
                    "rule_number": "number"
                  }
                ]
              ],
              "address": "string",
              "addresses": [
                "list",
                "string"
              ],
              "attached_to": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "availability_zones": [
                "list",
                "string"
              ],
              "cidrs": [
                "list",
                "string"
              ],
              "classic_load_balancer_listener": [
                "list",
                [
                  "object",
                  {
                    "instance_port": "number",
                    "load_balancer_port": "number"
                  }
                ]
              ],
              "component": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "customer_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "destination": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "destination_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "direction": "string",
              "elastic_load_balancer_listener": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "explanation_code": "string",
              "ingress_route_table": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "internet_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "load_balancer_arn": "string",
              "load_balancer_listener_port": "number",
              "load_balancer_target_group": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "load_balancer_target_groups": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "load_balancer_target_port": "number",
              "missing_component": "string",
              "nat_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "network_interface": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "packet_field": "string",
              "port": "number",
              "port_ranges": [
                "list",
                [
                  "object",
                  {
                    "from": "number",
                    "to": "number"
                  }
                ]
              ],
              "prefix_list": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "protocols": [
                "list",
                "string"
              ],
              "route_table": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "route_table_route": [
                "list",
                [
                  "object",
                  {
                    "destination_cidr": "string",
                    "destination_prefix_list_id": "string",
                    "egress_only_internet_gateway_id": "string",
                    "gateway_id": "string",
                    "instance_id": "string",
                    "nat_gateway_id": "string",
                    "network_interface_id": "string",
                    "origin": "string",
                    "transit_gateway_id": "string",
                    "vpc_peering_connection_id": "string"
                  }
                ]
              ],
              "security_group": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "security_group_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "direction": "string",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "prefix_list_id": "string",
                    "protocol": "string",
                    "security_group_id": "string"
                  }
                ]
              ],
              "security_groups": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "source_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "state": "string",
              "subnet": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "subnet_route_table": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway_attachment": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway_route_table": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway_route_table_route": [
                "list",
                [
                  "object",
                  {
                    "attachment_id": "string",
                    "destination_cidr": "string",
                    "prefix_list_id": "string",
                    "resource_id": "string",
                    "resource_type": "string",
                    "route_origin": "string",
                    "state": "string"
                  }
                ]
              ],
              "vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "vpc_endpoint": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "vpc_peering_connection": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "vpn_connection": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "vpn_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "filter_in_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "forward_path_components": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "egress": "bool",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "rule_action": "string",
                    "rule_number": "number"
                  }
                ]
              ],
              "additional_details": [
                "list",
                [
                  "object",
                  {
                    "additional_detail_type": "string",
                    "component": [
                      "list",
                      [
                        "object",
                        {
                          "arn": "string",
                          "id": "string",
                          "name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "attached_to": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "component": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "destination_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "inbound_header": [
                "list",
                [
                  "object",
                  {
                    "destination_addresses": [
                      "list",
                      "string"
                    ],
                    "destination_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "source_addresses": [
                      "list",
                      "string"
                    ],
                    "source_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "outbound_header": [
                "list",
                [
                  "object",
                  {
                    "destination_addresses": [
                      "list",
                      "string"
                    ],
                    "destination_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "source_addresses": [
                      "list",
                      "string"
                    ],
                    "source_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "route_table_route": [
                "list",
                [
                  "object",
                  {
                    "destination_cidr": "string",
                    "destination_prefix_list_id": "string",
                    "egress_only_internet_gateway_id": "string",
                    "gateway_id": "string",
                    "instance_id": "string",
                    "nat_gateway_id": "string",
                    "network_interface_id": "string",
                    "origin": "string",
                    "transit_gateway_id": "string",
                    "vpc_peering_connection_id": "string"
                  }
                ]
              ],
              "security_group_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "direction": "string",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "prefix_list_id": "string",
                    "protocol": "string",
                    "security_group_id": "string"
                  }
                ]
              ],
              "sequence_number": "number",
              "source_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "subnet": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway_route_table_route": [
                "list",
                [
                  "object",
                  {
                    "attachment_id": "string",
                    "destination_cidr": "string",
                    "prefix_list_id": "string",
                    "resource_id": "string",
                    "resource_type": "string",
                    "route_origin": "string",
                    "state": "string"
                  }
                ]
              ],
              "vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
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
        "optional": true,
        "type": "string"
      },
      "network_insights_analysis_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_insights_path_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "path_found": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "return_path_components": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "egress": "bool",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "rule_action": "string",
                    "rule_number": "number"
                  }
                ]
              ],
              "additional_details": [
                "list",
                [
                  "object",
                  {
                    "additional_detail_type": "string",
                    "component": [
                      "list",
                      [
                        "object",
                        {
                          "arn": "string",
                          "id": "string",
                          "name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "attached_to": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "component": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "destination_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "inbound_header": [
                "list",
                [
                  "object",
                  {
                    "destination_addresses": [
                      "list",
                      "string"
                    ],
                    "destination_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "source_addresses": [
                      "list",
                      "string"
                    ],
                    "source_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "outbound_header": [
                "list",
                [
                  "object",
                  {
                    "destination_addresses": [
                      "list",
                      "string"
                    ],
                    "destination_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "protocol": "string",
                    "source_addresses": [
                      "list",
                      "string"
                    ],
                    "source_port_ranges": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "route_table_route": [
                "list",
                [
                  "object",
                  {
                    "destination_cidr": "string",
                    "destination_prefix_list_id": "string",
                    "egress_only_internet_gateway_id": "string",
                    "gateway_id": "string",
                    "instance_id": "string",
                    "nat_gateway_id": "string",
                    "network_interface_id": "string",
                    "origin": "string",
                    "transit_gateway_id": "string",
                    "vpc_peering_connection_id": "string"
                  }
                ]
              ],
              "security_group_rule": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "direction": "string",
                    "port_range": [
                      "list",
                      [
                        "object",
                        {
                          "from": "number",
                          "to": "number"
                        }
                      ]
                    ],
                    "prefix_list_id": "string",
                    "protocol": "string",
                    "security_group_id": "string"
                  }
                ]
              ],
              "sequence_number": "number",
              "source_vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "subnet": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "transit_gateway_route_table_route": [
                "list",
                [
                  "object",
                  {
                    "attachment_id": "string",
                    "destination_cidr": "string",
                    "prefix_list_id": "string",
                    "resource_id": "string",
                    "resource_type": "string",
                    "route_origin": "string",
                    "state": "string"
                  }
                ]
              ],
              "vpc": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "start_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
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
      "warning_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2NetworkInsightsAnalysisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2NetworkInsightsAnalysis), &result)
	return &result
}
