package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "LineEnding",
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
	"inputEolSequence": "string",
	"outputEolSequence": "string",
	"inputCharset": "string",
	"outputCharset": "string",
	"postTransformationActionRenameAsExpression": "string"
},
*/

type stRouteStepLineEndingResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:LineEnding"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	PrecedingStep         types.String `tfsdk:"preceding_step"  helper:"precedingStep"`
	ConditionType         types.String `tfsdk:"condition_type"  helper:"conditionType,default:ALWAYS"`
	Condition             types.String `tfsdk:"condition"`
	ActionOnStepSuccess   types.String `tfsdk:"action_on_step_success"  helper:"actionOnStepSuccess,default:PROCEED"`
	ActionOnStepFailure   types.String `tfsdk:"action_on_step_failure"  helper:"actionOnStepFailure,default:FAIL"`
	Autostart             types.Bool   `tfsdk:"autostart"  helper:",default:false"`
	UsePrecedingStepFiles types.Bool   `tfsdk:"use_preceding_step_files"  helper:"usePrecedingStepFiles,default:false"`
	FileFilterExpression  types.String `tfsdk:"file_filter_expression"  helper:"fileFilterExpression"`
}

func init() {
	tfhelper.RegisterType("stRouteStepLineEndingResourceModel", &stRouteStepLineEndingResourceModel{})
}
