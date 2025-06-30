package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsChatbotSlackWorkspaceSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsChatbotSlackWorkspaceSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
