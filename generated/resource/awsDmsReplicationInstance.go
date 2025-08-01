package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsReplicationInstance = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allow_major_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "apply_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name_servers": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
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
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
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
      "replication_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_instance_private_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "replication_instance_public_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "replication_subnet_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "kerberos_authentication_settings": {
        "block": {
          "attributes": {
            "key_cache_secret_iam_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_cache_secret_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "krb5_file_contents": {
              "description_kind": "plain",
              "required": true,
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

func AwsDmsReplicationInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsReplicationInstance), &result)
	return &result
}
