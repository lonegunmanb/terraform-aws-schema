package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkmailDomain = `{
  "block": {
    "attributes": {
      "dkim_verification_status": {
        "computed": true,
        "description": "DKIM verification status for the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description": "Mail domain name to register.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "Whether this domain is the default mail domain for the organization.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_test_domain": {
        "computed": true,
        "description": "Whether this is the auto-provisioned test domain.",
        "description_kind": "plain",
        "type": "bool"
      },
      "organization_id": {
        "description": "Identifier of the WorkMail organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ownership_verification_status": {
        "computed": true,
        "description": "Domain ownership verification status.",
        "description_kind": "plain",
        "type": "string"
      },
      "records": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "hostname": "string",
              "type": "string",
              "value": "string"
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

func AwsWorkmailDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkmailDomain), &result)
	return &result
}
