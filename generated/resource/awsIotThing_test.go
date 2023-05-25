package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsIotThingSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsIotThingSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
