package generated_test

import (
	"testing"

	"github.com/lonegunmanb/terraform-modtm-schema/generated"
	"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
	assert.NotEmpty(t, generated.Resources)
	
}