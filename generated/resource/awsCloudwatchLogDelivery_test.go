package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsCloudwatchLogDeliverySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsCloudwatchLogDeliverySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}