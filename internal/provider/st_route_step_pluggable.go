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
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",default:Pluggable"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`

	// PrecedingStep types.String `tfsdk:"preceding_step"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	// Autostart             types.Bool `tfsdk:"autostart" helper:"autostart,default:false"`
	// UsePrecedingStepFiles types.Bool `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:false"`

	CustomProperties    types.Map `tfsdk:"custom_properties" helper:"customProperties,elementtype:string,optional,noread"`
	CustomPropertiesAll types.Map `tfsdk:"custom_properties_all" helper:"customProperties,elementtype:string,computed,nowrite,fieldMapOnRead:customProperties"`
}

func init() {
	tfhelper.RegisterType("stRouteStepPluggableResourceModel", &stRouteStepPluggableResourceModel{})
}
