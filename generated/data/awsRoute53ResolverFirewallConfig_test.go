package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsRoute53ResolverFirewallConfigSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsRoute53ResolverFirewallConfigSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
