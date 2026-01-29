package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsRamResourceShareAssociationsExclusiveSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsRamResourceShareAssociationsExclusiveSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
