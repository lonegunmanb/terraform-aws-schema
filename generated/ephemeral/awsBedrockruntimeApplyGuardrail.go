package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBedrockruntimeApplyGuardrail = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "action_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assessments": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "applied_guardrail_details": [
                "list",
                [
                  "object",
                  {
                    "guardrail_arn": "string",
                    "guardrail_id": "string",
                    "guardrail_origin": [
                      "list",
                      "string"
                    ],
                    "guardrail_ownership": "string",
                    "guardrail_version": "string"
                  }
                ]
              ],
              "automated_reasoning_policy": [
                "list",
                [
                  "object",
                  {
                    "findings": [
                      "list",
                      [
                        "object",
                        {
                          "impossible": [
                            "list",
                            [
                              "object",
                              {
                                "contradicting_rules": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "identifier": "string",
                                      "policy_version_arn": "string"
                                    }
                                  ]
                                ],
                                "logic_warning": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "type": "string"
                                    }
                                  ]
                                ],
                                "translation": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "confidence": "number",
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "invalid": [
                            "list",
                            [
                              "object",
                              {
                                "contradicting_rules": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "identifier": "string",
                                      "policy_version_arn": "string"
                                    }
                                  ]
                                ],
                                "logic_warning": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "type": "string"
                                    }
                                  ]
                                ],
                                "translation": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "confidence": "number",
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "no_translations": [
                            "list",
                            [
                              "object",
                              {}
                            ]
                          ],
                          "satisfiable": [
                            "list",
                            [
                              "object",
                              {
                                "claims_false_scenario": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "statements": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "claims_true_scenario": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "statements": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "logic_warning": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "type": "string"
                                    }
                                  ]
                                ],
                                "translation": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "confidence": "number",
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "too_complex": [
                            "list",
                            [
                              "object",
                              {}
                            ]
                          ],
                          "translation_ambiguous": [
                            "list",
                            [
                              "object",
                              {
                                "difference_scenarios": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "statements": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "options": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "translations": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "claims": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "logic": "string",
                                                  "natural_language": "string"
                                                }
                                              ]
                                            ],
                                            "confidence": "number",
                                            "premises": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "logic": "string",
                                                  "natural_language": "string"
                                                }
                                              ]
                                            ],
                                            "untranslated_claims": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "text": "string"
                                                }
                                              ]
                                            ],
                                            "untranslated_premises": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "text": "string"
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
                          "valid": [
                            "list",
                            [
                              "object",
                              {
                                "claims_true_scenario": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "statements": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "logic_warning": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "type": "string"
                                    }
                                  ]
                                ],
                                "supporting_rules": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "identifier": "string",
                                      "policy_version_arn": "string"
                                    }
                                  ]
                                ],
                                "translation": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "confidence": "number",
                                      "premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "logic": "string",
                                            "natural_language": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_claims": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
                                          }
                                        ]
                                      ],
                                      "untranslated_premises": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "text": "string"
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
                    ]
                  }
                ]
              ],
              "content_policy": [
                "list",
                [
                  "object",
                  {
                    "filters": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "confidence": "string",
                          "detected": "bool",
                          "filter_strength": "string",
                          "type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "contextual_grounding_policy": [
                "list",
                [
                  "object",
                  {
                    "filters": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "score": "number",
                          "threshold": "number",
                          "type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "invocation_metrics": [
                "list",
                [
                  "object",
                  {
                    "guardrail_coverage": [
                      "list",
                      [
                        "object",
                        {
                          "images": [
                            "list",
                            [
                              "object",
                              {
                                "guarded": "number",
                                "total": "number"
                              }
                            ]
                          ],
                          "text_characters": [
                            "list",
                            [
                              "object",
                              {
                                "guarded": "number",
                                "total": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "guardrail_processing_latency": "number",
                    "usage": [
                      "list",
                      [
                        "object",
                        {
                          "automated_reasoning_policies": "number",
                          "automated_reasoning_policy_units": "number",
                          "content_policy_image_units": "number",
                          "content_policy_units": "number",
                          "contextual_grounding_policy_units": "number",
                          "sensitive_information_policy_free_units": "number",
                          "sensitive_information_policy_units": "number",
                          "topic_policy_units": "number",
                          "word_policy_units": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "sensitive_information_policy": [
                "list",
                [
                  "object",
                  {
                    "pii_entities": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "match": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "regexes": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "match": "string",
                          "name": "string",
                          "regex": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "topic_policy": [
                "list",
                [
                  "object",
                  {
                    "topics": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "name": "string",
                          "type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "word_policy": [
                "list",
                [
                  "object",
                  {
                    "custom_words": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "match": "string"
                        }
                      ]
                    ],
                    "managed_word_lists": [
                      "list",
                      [
                        "object",
                        {
                          "action": "string",
                          "detected": "bool",
                          "match": "string",
                          "type": "string"
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
      "guardrail_coverage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "images": [
                "list",
                [
                  "object",
                  {
                    "guarded": "number",
                    "total": "number"
                  }
                ]
              ],
              "text_characters": [
                "list",
                [
                  "object",
                  {
                    "guarded": "number",
                    "total": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "guardrail_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "guardrail_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "text": "string"
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
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "usage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automated_reasoning_policies": "number",
              "automated_reasoning_policy_units": "number",
              "content_policy_image_units": "number",
              "content_policy_units": "number",
              "contextual_grounding_policy_units": "number",
              "sensitive_information_policy_free_units": "number",
              "sensitive_information_policy_units": "number",
              "topic_policy_units": "number",
              "word_policy_units": "number"
            }
          ]
        ]
      }
    },
    "block_types": {
      "content": {
        "block": {
          "block_types": {
            "text": {
              "block": {
                "attributes": {
                  "qualifiers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "text": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBedrockruntimeApplyGuardrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBedrockruntimeApplyGuardrail), &result)
	return &result
}
