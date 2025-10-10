package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOdbCloudAutonomousVmClusters = `{
  "block": {
    "attributes": {
      "cloud_autonomous_vm_clusters": {
        "computed": true,
        "description": "List of Cloud Autonomous VM Clusters. The list going to contain basic information about the cloud autonomous VM clusters.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "cloud_exadata_infrastructure_id": "string",
              "display_name": "string",
              "id": "string",
              "oci_resource_anchor_name": "string",
              "oci_url": "string",
              "ocid": "string",
              "odb_network_id": "string"
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

func AwsOdbCloudAutonomousVmClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOdbCloudAutonomousVmClusters), &result)
	return &result
}
