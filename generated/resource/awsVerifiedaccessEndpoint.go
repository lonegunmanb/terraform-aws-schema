package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVerifiedaccessEndpoint = `{
  "block": {
    "attributes": {
      "application_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attachment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "device_validation_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_certificate_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_domain_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_document": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "verified_access_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "verified_access_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cidr_options": {
        "block": {
          "attributes": {
            "cidr": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "to_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "load_balancer_options": {
        "block": {
          "attributes": {
            "load_balancer_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "to_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_interface_options": {
        "block": {
          "attributes": {
            "network_interface_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "port_range": {
              "block": {
                "attributes": {
                  "from_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "to_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rds_options": {
        "block": {
          "attributes": {
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rds_db_cluster_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rds_db_instance_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rds_db_proxy_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rds_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sse_specification": {
        "block": {
          "attributes": {
            "customer_managed_key_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVerifiedaccessEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVerifiedaccessEndpoint), &result)
	return &result
}
