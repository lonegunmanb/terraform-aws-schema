package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsRoute53ResolverFirewallDomainListSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsRoute53ResolverFirewallDomainListSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
