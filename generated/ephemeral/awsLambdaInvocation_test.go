package ephemeral_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/ephemeral"
	"github.com/stretchr/testify/assert"
)

func TestAwsLambdaInvocationSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := ephemeral.AwsLambdaInvocationSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
