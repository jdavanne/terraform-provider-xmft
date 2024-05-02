package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "Decompress",
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
	"filenameCollisionResolutionType": "OVERWRITE",
	"zipPassword": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepDecompressResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:Decompress"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepDecompressResourceModel", &stRouteStepDecompressResourceModel{})
}
