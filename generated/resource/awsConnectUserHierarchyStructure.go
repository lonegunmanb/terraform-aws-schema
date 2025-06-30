package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectUserHierarchyStructure = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
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
      "hierarchy_structure": {
        "block": {
          "block_types": {
            "level_five": {
              "block": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "level_four": {
              "block": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "level_one": {
              "block": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "level_three": {
              "block": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "level_two": {
              "block": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectUserHierarchyStructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectUserHierarchyStructure), &result)
	return &result
}
