package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontResponseHeadersPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cors_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_control_allow_credentials": "bool",
              "access_control_allow_headers": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "access_control_allow_methods": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "access_control_allow_origins": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "access_control_expose_headers": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "access_control_max_age_sec": "number",
              "origin_override": "bool"
            }
          ]
        ]
      },
      "custom_headers_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "items": [
                "set",
                [
                  "object",
                  {
                    "header": "string",
                    "override": "bool",
                    "value": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "etag": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remove_headers_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "items": [
                "set",
                [
                  "object",
                  {
                    "header": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "security_headers_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "content_security_policy": [
                "list",
                [
                  "object",
                  {
                    "content_security_policy": "string",
                    "override": "bool"
                  }
                ]
              ],
              "content_type_options": [
                "list",
                [
                  "object",
                  {
                    "override": "bool"
                  }
                ]
              ],
              "frame_options": [
                "list",
                [
                  "object",
                  {
                    "frame_option": "string",
                    "override": "bool"
                  }
                ]
              ],
              "referrer_policy": [
                "list",
                [
                  "object",
                  {
                    "override": "bool",
                    "referrer_policy": "string"
                  }
                ]
              ],
              "strict_transport_security": [
                "list",
                [
                  "object",
                  {
                    "access_control_max_age_sec": "number",
                    "include_subdomains": "bool",
                    "override": "bool",
                    "preload": "bool"
                  }
                ]
              ],
              "xss_protection": [
                "list",
                [
                  "object",
                  {
                    "mode_block": "bool",
                    "override": "bool",
                    "protection": "bool",
                    "report_uri": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "server_timing_headers_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "sampling_rate": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudfrontResponseHeadersPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontResponseHeadersPolicy), &result)
	return &result
}
