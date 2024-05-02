package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "PgpEncryption",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"fileFilterExpression": "string",
	"fileFilterExpressionType": "GLOB",
	"usePrecedingStepFiles": false,
	"encryptKeyOwnerExpression": "string",
	"encryptKeyExpression": "string",
	"encryptKeyOwnerExpressionType": "string",
	"encryptKeyExpressionType": "string",
	"signKeyOwnerExpression": "string",
	"signKeyExpression": "string",
	"signKeyOwnerExpressionType": "string",
	"signKeyExpressionType": "string",
	"useAsciiArmour": false,
	"compressionType": "string",
	"compressionLevel": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepPgpEncryptionResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:PgpEncryption"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPgpEncryptionResourceModel", &stRouteStepPgpEncryptionResourceModel{})
}
