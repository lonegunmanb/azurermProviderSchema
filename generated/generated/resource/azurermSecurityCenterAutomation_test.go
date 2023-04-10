package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurerm-provider-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSecurityCenterAutomationSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermSecurityCenterAutomationSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
