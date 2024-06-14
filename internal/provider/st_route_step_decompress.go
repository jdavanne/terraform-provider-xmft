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
	// PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,optional"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	// Autostart                                  types.Bool   `tfsdk:"autostart" helper:",default:false"`
	// UsePrecedingStepFiles                      types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,noread,nowrite"`
	FileFilterExpression                       types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType                   types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,default:GLOB"`
	FilenameCollisionResolutionType            types.String `tfsdk:"filename_collision_resolution_type" helper:"filenameCollisionResolutionType,enum:/RENAME_OLD/OVERWRITE/FAIL,emptyIsNull,default:OVERWRITE"`
	ZipPassword                                types.String `tfsdk:"zip_password" helper:"zipPassword,emptyIsNull,default:"`
	PostTransformationActionRenameAsExpression types.String `tfsdk:"post_transformation_action_rename_as_expression" helper:"postTransformationActionRenameAsExpression,emptyIsNull,default:"`
}

func init() {
	tfhelper.RegisterType("stRouteStepDecompressResourceModel", &stRouteStepDecompressResourceModel{})
}
