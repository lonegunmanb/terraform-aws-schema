package generated_test

import (
	"testing"

	"github.com/lonegunmanb/terraform-aws-schema/v5/generated"
	"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
	assert.NotEmpty(t, generated.Resources)
	assert.NotEmpty(t, generated.DataSources)
    assert.NotEmpty(t, generated.EphemeralResources)
}