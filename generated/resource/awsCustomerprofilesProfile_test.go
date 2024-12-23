package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsCustomerprofilesProfileSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsCustomerprofilesProfileSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
