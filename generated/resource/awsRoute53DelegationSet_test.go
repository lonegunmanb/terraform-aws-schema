package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsRoute53DelegationSetSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsRoute53DelegationSetSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}