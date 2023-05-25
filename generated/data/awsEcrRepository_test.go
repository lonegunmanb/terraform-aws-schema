package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsEcrRepositorySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsEcrRepositorySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
