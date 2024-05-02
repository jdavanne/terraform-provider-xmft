package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "Pluggable",
	"status": "ENABLED",
	"precedingStep": "string",
	"conditionType": "ALWAYS",
	"condition": "string",
	"actionOnStepSuccess": "PROCEED",
	"actionOnStepFailure": "FAIL",
	"autostart": false,
	"customProperties": {
		"property_1": "string",
		"property_2": "string"
	}
}
*/

type stRouteStepPluggableResourceModel struct {
	Id           types.String `tfsdk:"id" helper:",computed,state"`
	Type         types.String `tfsdk:"type" helper:",default:Pluggable"`
	Status       types.String `tfsdk:"status" helper:",default:ENABLED"`
	ExecuteRoute types.String `tfsdk:"execute_route" helper:"executeRoute,required"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPluggableResourceModel", &stRouteStepPluggableResourceModel{})
}
