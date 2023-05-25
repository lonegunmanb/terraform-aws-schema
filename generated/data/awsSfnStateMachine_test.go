package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsSfnStateMachineSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsSfnStateMachineSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
