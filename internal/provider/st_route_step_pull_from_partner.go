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
	Id           types.String `tfsdk:"id" helper:",computed,state"`
	Type         types.String `tfsdk:"type" helper:",required"` // ExecuteRoute
	Status       types.String `tfsdk:"status" helper:",default:ENABLED"`
	ExecuteRoute types.String `tfsdk:"execute_route" helper:"executeRoute,required"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPullFromPartnerResourceModel", &stRouteStepPullFromPartnerResourceModel{})
}
