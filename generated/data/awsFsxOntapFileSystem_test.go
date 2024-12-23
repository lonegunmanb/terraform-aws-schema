package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsFsxOntapFileSystemSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsFsxOntapFileSystemSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
