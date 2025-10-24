package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbNetwork = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The name of the Availability Zone (AZ) where the odb network is located. Changing this will force terraform to create new resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "description": "The AZ ID of the AZ where the ODB network is located. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_subnet_cidr": {
        "description": "The CIDR range of the backup subnet for the ODB network. Changing this will force terraform to create new resource.\n\tConstraints:\n\t   - Must not overlap with the CIDR range of the client subnet.\n\t   - Must not overlap with the CIDR ranges of the VPCs that are connected to the\n\t   ODB network.\n\t   - Must not use the following CIDR ranges that are reserved by OCI:\n\t   - 100.106.0.0/16 and 100.107.0.0/16\n\t   - 169.254.0.0/16\n\t   - 224.0.0.0 - 239.255.255.255\n\t   - 240.0.0.0 - 255.255.255.255",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_subnet_cidr": {
        "description": "The CIDR notation for the network resource. Changing this will force terraform to create new resource.\n Constraints:\n  \t - Must not overlap with the CIDR range of the backup subnet.\n   \t- Must not overlap with the CIDR ranges of the VPCs that are connected to the\n   ODB network.\n  \t- Must not use the following CIDR ranges that are reserved by OCI:\n  \t - 100.106.0.0/16 and 100.107.0.0/16\n  \t - 169.254.0.0/16\n   \t- 224.0.0.0 - 239.255.255.255\n   \t- 240.0.0.0 - 255.255.255.255",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the ODB network was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_name": {
        "description": "The name of the custom domain that the network is located. custom_domain_name and default_dns_prefix both can't be given.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_dns_prefix": {
        "description": "The default DNS prefix for the network resource. Changing this will force terraform to create new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_associated_resources": {
        "computed": true,
        "description": "If set to true deletes associated OCI resources. Default false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The user-friendly name for the odb network. Changing this will force terraform to create a new resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "managed_services": {
        "computed": true,
        "description": "The managed services configuration for the ODB network.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "managed_s3_backup_access": [
                "list",
                [
                  "object",
                  {
                    "ipv4_addresses": [
                      "set",
                      "string"
                    ],
                    "status": "string"
                  }
                ]
              ],
              "managed_service_ipv4_cidrs": [
                "set",
                "string"
              ],
              "resource_gateway_arn": "string",
              "s3_access": [
                "list",
                [
                  "object",
                  {
                    "domain_name": "string",
                    "ipv4_addresses": [
                      "set",
                      "string"
                    ],
                    "s3_policy_document": "string",
                    "status": "string"
                  }
                ]
              ],
              "service_network_arn": "string",
              "service_network_endpoint": [
                "list",
                [
                  "object",
                  {
                    "vpc_endpoint_id": "string",
                    "vpc_endpoint_type": "string"
                  }
                ]
              ],
              "zero_etl_access": [
                "list",
                [
                  "object",
                  {
                    "cidr": "string",
                    "status": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "oci_dns_forwarding_configs": {
        "computed": true,
        "description": "The DNS resolver endpoint in OCI for forwarding DNS queries for the ociPrivateZone domain.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain_name": "string",
              "oci_dns_listener_ip": "string"
            }
          ]
        ]
      },
      "oci_network_anchor_id": {
        "computed": true,
        "description": "The unique identifier of the OCI network anchor for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_network_anchor_url": {
        "computed": true,
        "description": "The URL of the OCI network anchor for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_vcn_id": {
        "computed": true,
        "description": "The unique identifier  Oracle Cloud ID (OCID) of the OCI VCN for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_vcn_url": {
        "computed": true,
        "description": "The URL of the OCI VCN for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "peered_cidrs": {
        "computed": true,
        "description": "The list of CIDR ranges from the peered VPC that are allowed access to the ODB network. Please refer odb network peering documentation.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "percent_progress": {
        "computed": true,
        "description": "The amount of progress made on the current operation on the ODB network, expressed as a percentage.",
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_access": {
        "description": "Specifies the configuration for Amazon S3 access from the ODB network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_policy_document": {
        "description": "Specifies the endpoint policy for Amazon S3 access from the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the network resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Additional information about the current status of the ODB network.",
        "description_kind": "plain",
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
        "type": [
          "map",
          "string"
        ]
      },
      "zero_etl_access": {
        "description": "Specifies the configuration for Zero-ETL access from the ODB network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsOdbNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbNetwork), &result)
	return &result
}
