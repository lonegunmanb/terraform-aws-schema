package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGuarddutyOrganizationConfiguration = `{
  "block": {
    "attributes": {
      "auto_enable_organization_members": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "detector_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "datasources": {
        "block": {
          "block_types": {
            "kubernetes": {
              "block": {
                "block_types": {
                  "audit_logs": {
                    "block": {
                      "attributes": {
                        "enable": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "malware_protection": {
              "block": {
                "block_types": {
                  "scan_ec2_instance_with_findings": {
                    "block": {
                      "block_types": {
                        "ebs_volumes": {
                          "block": {
                            "attributes": {
                              "auto_enable": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
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
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "s3_logs": {
              "block": {
                "attributes": {
                  "auto_enable": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGuarddutyOrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGuarddutyOrganizationConfiguration), &result)
	return &result
}
