package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsBedrockFoundationModelSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsBedrockFoundationModelSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
