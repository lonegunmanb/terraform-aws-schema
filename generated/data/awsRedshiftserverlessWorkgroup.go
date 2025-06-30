package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftserverlessWorkgroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "port": "number",
              "vpc_endpoint": [
                "list",
                [
                  "object",
                  {
                    "network_interface": [
                      "list",
                      [
                        "object",
                        {
                          "availability_zone": "string",
                          "network_interface_id": "string",
                          "private_ip_address": "string",
                          "subnet_id": "string"
                        }
                      ]
                    ],
                    "vpc_endpoint_id": "string",
                    "vpc_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
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
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "track_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workgroup_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workgroup_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftserverlessWorkgroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftserverlessWorkgroup), &result)
	return &result
}
