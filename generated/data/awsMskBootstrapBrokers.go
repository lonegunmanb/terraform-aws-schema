package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskBootstrapBrokers = `{
  "block": {
    "attributes": {
      "bootstrap_brokers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMskBootstrapBrokersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskBootstrapBrokers), &result)
	return &result
}
