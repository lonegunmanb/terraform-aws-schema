package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsCodecommitRepositorySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsCodecommitRepositorySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
