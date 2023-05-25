package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsElasticBeanstalkApplicationSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsElasticBeanstalkApplicationSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
