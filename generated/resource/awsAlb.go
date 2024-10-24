package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAlb = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn_suffix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_keep_alive": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "customer_owned_ipv4_pool": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desync_mitigation_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_record_client_routing_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "drop_invalid_header_fields": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_cross_zone_load_balancing": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_http2": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_tls_version_and_cipher_suite_headers": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_waf_fail_open": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_xff_client_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_zonal_shift": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enforce_security_group_inbound_rules_on_private_link_traffic": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "idle_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "internal": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preserve_host_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "subnets": {
        "computed": true,
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "xff_header_processing_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "access_logs": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "prefix": {
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
      "connection_logs": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "prefix": {
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
      "subnet_mapping": {
        "block": {
          "attributes": {
            "allocation_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outpost_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_ipv4_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AwsAlbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAlb), &result)
	return &result
}
