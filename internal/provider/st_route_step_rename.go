package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
	"type": "Rename",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"usePrecedingStepFiles": false,
	"outputFileName": "string",
	"fileFilterExpression": "string",
	"fileFilterExpressionType": "GLOB"
},
*/

type stRouteStepRenameResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:Rename"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,computed,nowrite"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:/ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",optional"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	Autostart                types.Bool   `tfsdk:"autostart" helper:"autostart,default:false"`
	UsePrecedingStepFiles    types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:true"`
	OutputFileName           types.String `tfsdk:"output_file_name" helper:"outputFileName,required"`
	FileFilterExpression     types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
}

func init() {
	tfhelper.RegisterType("stRouteStepRenameResourceModel", &stRouteStepRenameResourceModel{})
}
