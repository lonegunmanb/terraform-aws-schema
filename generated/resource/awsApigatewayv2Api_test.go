package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsApigatewayv2ApiSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsApigatewayv2ApiSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}