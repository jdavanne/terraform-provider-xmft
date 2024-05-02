package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "CharactersReplace",
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
	"findCharacterSequence": "string",
	"lineStrip": "string",
	"replaceCharacterSequence": "string",
	"inputCharset": "string",
	"outputCharset": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepCharacterReplaceResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:CharactersReplace"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepCharacterReplaceResourceModel", &stRouteStepCharacterReplaceResourceModel{})
}
