package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockFoundationModelAgreementOffers = `{
  "block": {
    "attributes": {
      "model_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "offer_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "offers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "offer_id": "string",
              "offer_token": "string",
              "term_details": [
                "list",
                [
                  "object",
                  {
                    "legal_term": [
                      "list",
                      [
                        "object",
                        {
                          "url": "string"
                        }
                      ]
                    ],
                    "support_term": [
                      "list",
                      [
                        "object",
                        {
                          "refund_policy_description": "string"
                        }
                      ]
                    ],
                    "usage_based_pricing_term": [
                      "list",
                      [
                        "object",
                        {
                          "rate_card": [
                            "list",
                            [
                              "object",
                              {
                                "description": "string",
                                "dimension": "string",
                                "price": "string",
                                "unit": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "validity_term": [
                      "list",
                      [
                        "object",
                        {
                          "agreement_duration": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockFoundationModelAgreementOffersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockFoundationModelAgreementOffers), &result)
	return &result
}
