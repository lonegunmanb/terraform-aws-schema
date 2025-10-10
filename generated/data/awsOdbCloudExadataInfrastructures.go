package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudExadataInfrastructures = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructures": {
        "computed": true,
        "description": "List of Cloud Exadata Infrastructures. Returns basic information about the Cloud Exadata Infrastructures.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "display_name": "string",
              "id": "string",
              "oci_resource_anchor_name": "string",
              "oci_url": "string",
              "ocid": "string"
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

func AwsOdbCloudExadataInfrastructuresSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudExadataInfrastructures), &result)
	return &result
}
