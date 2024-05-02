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
	Id           types.String `tfsdk:"id" helper:",computed,state"`
	Type         types.String `tfsdk:"type" helper:",required"` // Publish
	Status       types.String `tfsdk:"status" helper:",default:ENABLED"`
	ExecuteRoute types.String `tfsdk:"execute_route" helper:"executeRoute,required"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPublishToAccountResourceModel", &stRouteStepPublishToAccountResourceModel{})
}
