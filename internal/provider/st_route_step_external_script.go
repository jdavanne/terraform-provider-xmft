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

	// PrecedingStep       types.String `tfsdk:"preceding_step" helper:"precedingStep,optional"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`
	// Autostart           types.Bool   `tfsdk:"autostart" helper:",default:false"`
	ScriptPath    types.String `tfsdk:"script_path" helper:"scriptPath,required"`
	LogOutput     types.Bool   `tfsdk:"log_output" helper:"logOutput,default:false"`
	RootExecution types.Bool   `tfsdk:"root_execution" helper:"rootExecution,default:false"`
}

func init() {
	tfhelper.RegisterType("stRouteStepExternalScriptResourceModel", &stRouteStepExternalScriptResourceModel{})
}
