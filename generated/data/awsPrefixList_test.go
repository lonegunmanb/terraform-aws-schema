package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsPrefixListSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsPrefixListSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
