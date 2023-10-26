package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-modtm-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestModtmTelemetrySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.ModtmTelemetrySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
