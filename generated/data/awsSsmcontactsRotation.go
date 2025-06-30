package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmcontactsRotation = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "recurrence": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "daily_settings": [
                "list",
                [
                  "object",
                  {
                    "hour_of_day": "number",
                    "minute_of_hour": "number"
                  }
                ]
              ],
              "monthly_settings": [
                "list",
                [
                  "object",
                  {
                    "day_of_month": "number",
                    "hand_off_time": [
                      "list",
                      [
                        "object",
                        {
                          "hour_of_day": "number",
                          "minute_of_hour": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "number_of_on_calls": "number",
              "recurrence_multiplier": "number",
              "shift_coverages": [
                "list",
                [
                  "object",
                  {
                    "coverage_times": [
                      "list",
                      [
                        "object",
                        {
                          "end": [
                            "list",
                            [
                              "object",
                              {
                                "hour_of_day": "number",
                                "minute_of_hour": "number"
                              }
                            ]
                          ],
                          "start": [
                            "list",
                            [
                              "object",
                              {
                                "hour_of_day": "number",
                                "minute_of_hour": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "map_block_key": "string"
                  }
                ]
              ],
              "weekly_settings": [
                "list",
                [
                  "object",
                  {
                    "day_of_week": "string",
                    "hand_off_time": [
                      "list",
                      [
                        "object",
                        {
                          "hour_of_day": "number",
                          "minute_of_hour": "number"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
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
      },
      "time_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmcontactsRotationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmcontactsRotation), &result)
	return &result
}
