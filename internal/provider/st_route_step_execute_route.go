package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"type": "ExecuteRoute",
	"status": "ENABLED",
	"precedingStep": "string",
	"executeRoute": "string"
}
*/

type stRouteStepExecuteRouteResourceModel struct {
	// Id           types.String `tfsdk:"id" helper:",computed,state"`
	Type         types.String `tfsdk:"type" helper:",default:ExecuteRoute"`
	Status       types.String `tfsdk:"status" helper:",default:ENABLED"`
	ExecuteRoute types.String `tfsdk:"execute_route" helper:"executeRoute,required"`
}

func init() {
	tfhelper.RegisterType("stRouteStepExecuteRouteResourceModel", &stRouteStepExecuteRouteResourceModel{})
}
