package ephemeral_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v4/generated/ephemeral"
	"github.com/stretchr/testify/assert"
)

func TestAzurermKeyVaultCertificateSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := ephemeral.AzurermKeyVaultCertificateSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}