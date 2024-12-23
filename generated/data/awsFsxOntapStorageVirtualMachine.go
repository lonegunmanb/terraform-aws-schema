package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxOntapStorageVirtualMachine = `{
  "block": {
    "attributes": {
      "active_directory_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "netbios_name": "string",
              "self_managed_active_directory_configuration": [
                "list",
                [
                  "object",
                  {
                    "dns_ips": [
                      "set",
                      "string"
                    ],
                    "domain_name": "string",
                    "file_system_administrators_group": "string",
                    "organizational_unit_distinguished_name": "string",
                    "username": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "iscsi": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "nfs": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "smb": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
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
      "file_system_id": {
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
      "lifecycle_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_transition_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "message": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subtype": {
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
      "uuid": {
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
                "list",
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

func AwsFsxOntapStorageVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxOntapStorageVirtualMachine), &result)
	return &result
}
