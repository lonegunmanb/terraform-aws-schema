package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBackupVaultLockConfiguration = `{
  "block": {
    "attributes": {
      "backup_vault_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "changeable_for_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_retention_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_retention_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AwsBackupVaultLockConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBackupVaultLockConfiguration), &result)
	return &result
}
