package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStoragegatewayCachedIscsiVolume = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "chap_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "gateway_arn": {
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
      "kms_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lun_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "network_interface_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_volume_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_size_in_bytes": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsStoragegatewayCachedIscsiVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsStoragegatewayCachedIscsiVolume), &result)
	return &result
}
