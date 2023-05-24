package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsMacie2AccountSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsMacie2AccountSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
