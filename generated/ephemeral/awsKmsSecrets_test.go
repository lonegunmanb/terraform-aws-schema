package ephemeral_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/ephemeral"
	"github.com/stretchr/testify/assert"
)

func TestAwsKmsSecretsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := ephemeral.AwsKmsSecretsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
