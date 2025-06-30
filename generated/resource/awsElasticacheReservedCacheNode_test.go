package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsElasticacheReservedCacheNodeSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsElasticacheReservedCacheNodeSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
