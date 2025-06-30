package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDbProxyTarget = `{
  "block": {
    "attributes": {
      "db_cluster_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_proxy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
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
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "rds_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tracked_cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDbProxyTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbProxyTarget), &result)
	return &result
}
