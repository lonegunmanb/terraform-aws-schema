package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsDatasyncAgentSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsDatasyncAgentSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
