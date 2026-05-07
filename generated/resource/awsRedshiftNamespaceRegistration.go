package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftNamespaceRegistration = `{
  "block": {
    "attributes": {
      "consumer_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioned_cluster_identifier": {
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
      },
      "serverless_namespace_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serverless_workgroup_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftNamespaceRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftNamespaceRegistration), &result)
	return &result
}
