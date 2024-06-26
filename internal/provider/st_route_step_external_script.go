package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "ExternalScript",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"scriptPath": "string",
	"logOutput": false,
	"rootExecution": false
},
*/

type stRouteStepExternalScriptResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:ExternalScript"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
}

func init() {
	tfhelper.RegisterType("stRouteStepExternalScriptResourceModel", &stRouteStepExternalScriptResourceModel{})
}
