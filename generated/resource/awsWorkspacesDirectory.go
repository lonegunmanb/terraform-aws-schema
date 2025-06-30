package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspacesDirectory = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iam_role_id": {
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
      "ip_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registration_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
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
      "user_identity_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_directory_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_directory_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "active_directory_config": {
        "block": {
          "attributes": {
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_account_secret_arn": {
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
      "certificate_based_auth_properties": {
        "block": {
          "attributes": {
            "certificate_authority_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "computed": true,
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
      "saml_properties": {
        "block": {
          "attributes": {
            "relay_state_parameter_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_access_url": {
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
      "self_service_permissions": {
        "block": {
          "attributes": {
            "change_compute_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "increase_volume_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rebuild_workspace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "restart_workspace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "switch_running_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "workspace_access_properties": {
        "block": {
          "attributes": {
            "device_type_android": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_chromeos": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_ios": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_linux": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_osx": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_web": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_windows": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type_zeroclient": {
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
      "workspace_creation_properties": {
        "block": {
          "attributes": {
            "custom_security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_ou": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_internet_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_maintenance_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "user_enabled_as_local_administrator": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWorkspacesDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspacesDirectory), &result)
	return &result
}
