package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "PullFromPartner",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"targetAccountExpression": "string",
	"targetAccountExpressionType": "string",
	"transferSiteExpression": "string",
	"transferSiteExpressionType": "LIST",
	"remoteFolderPathExpression": "string",
	"remoteFolderPathExpressionType": "SIMPLE",
	"remoteFileNameExpression": "string",
	"remoteFileNameExpressionType": "GLOB",
	"localFolderPathExpression": "string",
	"localFolderPathExpressionType": "SIMPLE",
	"localFileNameExpression": "string",
	"localFileNameExpressionType": "SIMPLE"
}
*/

type stRouteStepPullFromPartnerResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:PullFromPartner"` // ExecuteRoute
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,computed,nowrite"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:/ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",optional"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	Autostart                      types.Bool   `tfsdk:"autostart" helper:"autostart,default:false"`
	TargetAccountExpression        types.String `tfsdk:"target_account_expression" helper:"targetAccountExpression,emptyIsNull,required"`
	TargetAccountExpressionType    types.String `tfsdk:"target_account_expression_type" helper:"targetAccountExpressionType,enum:/NAME/EXPRESSION,default:NAME"`
	TransferSiteExpression         types.String `tfsdk:"transfer_site_expression" helper:"transferSiteExpression,required"`
	TransferSiteExpressionType     types.String `tfsdk:"transfer_site_expression_type" helper:"transferSiteExpressionType,enum:/LIST,default:LIST"`
	RemoteFolderPathExpression     types.String `tfsdk:"remote_folder_path_expression" helper:"remoteFolderPathExpression,default:"`
	RemoteFolderPathExpressionType types.String `tfsdk:"remote_folder_path_expression_type" helper:"remoteFolderPathExpressionType,enum:/SIMPLE/EXPRESSION,default:SIMPLE"`
	RemoteFileNameExpression       types.String `tfsdk:"remote_file_name_expression" helper:"remoteFileNameExpression,optional"`
	RemoteFileNameExpressionType   types.String `tfsdk:"remote_file_name_expression_type" helper:"remoteFileNameExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
	LocalFolderPathExpression      types.String `tfsdk:"local_folder_path_expression" helper:"localFolderPathExpression,optional"`
	LocalFolderPathExpressionType  types.String `tfsdk:"local_folder_path_expression_type" helper:"localFolderPathExpressionType,enum:/SIMPLE/EXPRESSION,default:SIMPLE"`
	LocalFileNameExpression        types.String `tfsdk:"local_file_name_expression" helper:"localFileNameExpression,optional"`
	LocalFileNameExpressionType    types.String `tfsdk:"local_file_name_expression_type" helper:"localFileNameExpressionType,enum:/SIMPLE/EXPRESSION,default:SIMPLE"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPullFromPartnerResourceModel", &stRouteStepPullFromPartnerResourceModel{})
}
