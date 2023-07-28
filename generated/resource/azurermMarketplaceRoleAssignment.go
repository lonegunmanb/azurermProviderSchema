package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMarketplaceRoleAssignment = `{
  "block": {
    "attributes": {
      "condition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "condition_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delegated_managed_identity_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_definition_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_definition_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_service_principal_aad_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMarketplaceRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMarketplaceRoleAssignment), &result)
	return &result
}
