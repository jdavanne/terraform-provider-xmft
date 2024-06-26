package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "PgpDecryption",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"usePrecedingStepFiles": false,
	"fileFilterExpression": "string",
	"fileFilterExpressionType": "GLOB",
	"requireEncryption": false,
	"requireTrustedSignature": false,
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepPgpDecryptionResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:PgpDecryption"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPgpDecryptionResourceModel", &stRouteStepPgpDecryptionResourceModel{})
}
