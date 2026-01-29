package data

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
        "description": "The availability zone where the ODB network is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the AZ where the ODB network is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_subnet_cidr": {
        "computed": true,
        "description": " The CIDR range of the backup subnet for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_subnet_cidr": {
        "computed": true,
        "description": "The CIDR notation for the network resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the ODB network was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_name": {
        "computed": true,
        "description": "The name of the custom domain that the network is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_dns_prefix": {
        "computed": true,
        "description": "The default DNS prefix for the network resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Display name for the network resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
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
              "kms_access": [
                "list",
                [
                  "object",
                  {
                    "domain_name": "string",
                    "ipv4_addresses": [
                      "list",
                      "string"
                    ],
                    "kms_policy_document": "string",
                    "status": "string"
                  }
                ]
              ],
              "managed_s3_backup_access": [
                "list",
                [
                  "object",
                  {
                    "ipv4_addresses": [
                      "list",
                      "string"
                    ],
                    "status": "string"
                  }
                ]
              ],
              "managed_service_ipv4_cidrs": [
                "list",
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
                      "list",
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
              "sts_access": [
                "list",
                [
                  "object",
                  {
                    "domain_name": "string",
                    "ipv4_addresses": [
                      "list",
                      "string"
                    ],
                    "status": "string",
                    "sts_policy_document": "string"
                  }
                ]
              ],
              "zero_tl_access": [
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
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
