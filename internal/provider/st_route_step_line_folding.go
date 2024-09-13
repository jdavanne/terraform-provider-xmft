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

	PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,computed,nowrite"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:/ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",optional"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	Autostart                                  types.Bool   `tfsdk:"autostart" helper:"autostart,default:false"`
	UsePrecedingStepFiles                      types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:true"`
	FileFilterExpression                       types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType                   types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
	FileFoldWidth                              types.Int64  `tfsdk:"file_fold_width" helper:"fileFoldWidth,default:1"`
	InputCharset                               types.String `tfsdk:"input_charset" helper:"inputCharset,default:UTF-8"`
	OutputCharset                              types.String `tfsdk:"output_charset" helper:"outputCharset,optional"`
	PostTransformationActionRenameAsExpression types.String `tfsdk:"post_transformation_action_rename_as_expression" helper:"postTransformationActionRenameAsExpression,optional"`
}

func init() {
	tfhelper.RegisterType("stRouteStepLineFoldingResourceModel", &stRouteStepLineFoldingResourceModel{})
}
