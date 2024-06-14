package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "Publish",
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
	"filenameCollisionResolutionType": "string",
	"targetAccountExpression": "string",
	"targetAccountExpressionType": "string",
	"targetFolderExpression": "string",
	"targetFolderExpressionType": "string",
	"publishFileAs": "string",
	"triggerSubscription": false,
	"postRoutingActionType": "string",
	"disableAutoCreateTargetFolder": false,
	"postRoutingActionRenameExpression": "string"
},
*/

type stRouteStepPublishToAccountResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:Publish"` // Publish
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	// PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,optional"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	// Autostart                                  types.Bool   `tfsdk:"autostart" helper:",default:false"`
	// UsePrecedingStepFiles                      types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,noread,nowrite"`
	FileFilterExpression              types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType          types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
	FilenameCollisionResolutionType   types.String `tfsdk:"filename_collision_resolution_type" helper:"filenameCollisionResolutionType,enum:/RENAME_OLD/OVERWRITE/FAIL,emptyIsNull,default:OVERWRITE"`
	TargetAccountExpression           types.String `tfsdk:"target_account_expression" helper:"targetAccountExpression,required"`
	TargetAccountExpressionType       types.String `tfsdk:"target_account_expression_type" helper:"targetAccountExpressionType,enum:/NAME/EXPRESSION,emptyIsNull,default:NAME"`
	TargetFolderExpression            types.String `tfsdk:"target_folder_expression" helper:"targetFolderExpression,emptyIsNull,default:/"`
	TargetFolderExpressionType        types.String `tfsdk:"target_folder_expression_type" helper:"targetFolderExpressionType,enum:/SIMPLE/EXPRESSION,emptyIsNull,default:SIMPLE"`
	PublishFileAs                     types.String `tfsdk:"publish_file_as" helper:"publishFileAs,emptyIsNull,default:"`
	TriggerSubscription               types.Bool   `tfsdk:"trigger_subscription" helper:"triggerSubscription,default:false"`
	PostRoutingActionType             types.String `tfsdk:"post_routing_action_type" helper:"postRoutingActionType,emptyIsNull,default:d"`
	DisableAutoCreateTargetFolder     types.Bool   `tfsdk:"disable_auto_create_target_folder" helper:"disableAutoCreateTargetFolder,default:false"`
	PostRoutingActionRenameExpression types.String `tfsdk:"post_routing_action_rename_expression" helper:"postRoutingActionRenameExpression,emptyIsNull,default:"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPublishToAccountResourceModel", &stRouteStepPublishToAccountResourceModel{})
}
