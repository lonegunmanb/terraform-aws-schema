package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsS3AccessPointSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsS3AccessPointSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
