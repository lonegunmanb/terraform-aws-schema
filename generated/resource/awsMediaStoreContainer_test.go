package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v6/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsMediaStoreContainerSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsMediaStoreContainerSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
