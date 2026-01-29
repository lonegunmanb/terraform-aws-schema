package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2ManagedRuleGroup = `{
  "block": {
    "attributes": {
      "available_labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "consumed_labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "label_namespace": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": [
                "list",
                [
                  "object",
                  {
                    "allow": [
                      "list",
                      [
                        "object",
                        {
                          "custom_request_handling": [
                            "list",
                            [
                              "object",
                              {
                                "insert_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "block": [
                      "list",
                      [
                        "object",
                        {
                          "custom_response": [
                            "list",
                            [
                              "object",
                              {
                                "custom_response_body_key": "string",
                                "response_code": "number",
                                "response_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "captcha": [
                      "list",
                      [
                        "object",
                        {
                          "custom_request_handling": [
                            "list",
                            [
                              "object",
                              {
                                "insert_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "challenge": [
                      "list",
                      [
                        "object",
                        {
                          "custom_request_handling": [
                            "list",
                            [
                              "object",
                              {
                                "insert_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "count": [
                      "list",
                      [
                        "object",
                        {
                          "custom_request_handling": [
                            "list",
                            [
                              "object",
                              {
                                "insert_header": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "name": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "name": "string"
            }
          ]
        ]
      },
      "scope": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vendor_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWafv2ManagedRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2ManagedRuleGroup), &result)
	return &result
}
