package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "LineFolding",
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
	"fileFoldWidth": 1,
	"inputCharset": "string",
	"outputCharset": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepLineFoldingResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:LineFolding"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepLineFoldingResourceModel", &stRouteStepLineFoldingResourceModel{})
}
