package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatazonePolicyGrant = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "grant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
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
      }
    },
    "block_types": {
      "detail": {
        "block": {
          "block_types": {
            "add_to_project_member_pool": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_asset_type": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_domain_unit": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_environment": {
              "block": {
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_environment_from_blueprint": {
              "block": {
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_environment_profile": {
              "block": {
                "attributes": {
                  "domain_unit_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_form_type": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_glossary": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_project": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "create_project_from_project_profile": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "project_profiles": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "delegate_create_environment_profile": {
              "block": {
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "override_domain_unit_owners": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "override_project_owners": {
              "block": {
                "attributes": {
                  "include_child_domain_units": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "use_asset_type": {
              "block": {
                "attributes": {
                  "domain_unit_id": {
                    "description_kind": "plain",
                    "optional": true,
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
      },
      "principal": {
        "block": {
          "block_types": {
            "domain_unit": {
              "block": {
                "attributes": {
                  "domain_unit_designation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "domain_unit_identifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "all_domain_units_grant_filter": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "group": {
              "block": {
                "attributes": {
                  "group_identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "project": {
              "block": {
                "attributes": {
                  "project_designation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "project_identifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "domain_unit_filter": {
                    "block": {
                      "attributes": {
                        "domain_unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "include_child_domain_units": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
            },
            "user": {
              "block": {
                "attributes": {
                  "user_identifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "all_users_grant_filter": {
                    "block": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDatazonePolicyGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatazonePolicyGrant), &result)
	return &result
}
