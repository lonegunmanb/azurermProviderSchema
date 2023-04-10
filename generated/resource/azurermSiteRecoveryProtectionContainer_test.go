package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSiteRecoveryProtectionContainerSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermSiteRecoveryProtectionContainerSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
