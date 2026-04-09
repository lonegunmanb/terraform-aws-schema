package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3FilesMountTarget = `{
  "block": {
    "attributes": {
      "availability_zone_id": {
        "computed": true,
        "description": "Availability Zone ID",
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_id": {
        "computed": true,
        "description": "File system ID",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Mount target ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv4_address": {
        "computed": true,
        "description": "IPv4 address",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_address": {
        "computed": true,
        "description": "IPv6 address",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "Network interface ID",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AWS account ID of the owner",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "Security group IDs",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "Mount target status",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "Status message",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "Subnet ID",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "VPC ID",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3FilesMountTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3FilesMountTarget), &result)
	return &result
}
