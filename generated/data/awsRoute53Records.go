package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53Records = `{
  "block": {
    "attributes": {
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_record_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alias_target": [
                "object",
                {
                  "dns_name": "string",
                  "evaluate_target_health": "bool",
                  "hosted_zone_id": "string"
                }
              ],
              "cidr_routing_config": [
                "object",
                {
                  "collection_id": "string",
                  "location_name": "string"
                }
              ],
              "failover": "string",
              "geolocation": [
                "object",
                {
                  "continent_code": "string",
                  "country_code": "string",
                  "subdivision_code": "string"
                }
              ],
              "geoproximity_location": [
                "object",
                {
                  "aws_region": "string",
                  "bias": "number",
                  "coordinates": [
                    "object",
                    {
                      "latitude": "string",
                      "longitude": "string"
                    }
                  ],
                  "local_zone_group": "string"
                }
              ],
              "health_check_id": "string",
              "multi_value_answer": "bool",
              "name": "string",
              "region": "string",
              "resource_records": [
                "list",
                [
                  "object",
                  {
                    "value": "string"
                  }
                ]
              ],
              "set_identifier": "string",
              "traffic_policy_instance_id": "string",
              "ttl": "number",
              "type": "string",
              "weight": "number"
            }
          ]
        ]
      },
      "zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53RecordsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53Records), &result)
	return &result
}
