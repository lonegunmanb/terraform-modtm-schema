package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-modtm-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestModtmModuleSourceSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.ModtmModuleSourceSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
