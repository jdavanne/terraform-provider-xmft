package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stRouteSimpleResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Description types.String `tfsdk:"description" helper:",default"`
	Type        types.String `tfsdk:"type" helper:",default:SIMPLE"`
	// ManagedByCG   types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,optional,computed"`
	// RouteTemplate types.String `tfsdk:"route_template" helper:"routeTemplate,required"`
	// Account                     types.String   `tfsdk:"account" helper:",optional"`
	// Subscriptions               []types.String `tfsdk:"subscriptions"`
	ConditionType               types.String `tfsdk:"condition_type" helper:"conditionType,default:ALWAYS"`
	Condition                   types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	FailureEmailNotification    types.Bool   `tfsdk:"failure_email_notification" helper:"failureEmailNotification,emptyIsNull,default:"`
	FailureEmailTemplate        types.String `tfsdk:"failure_email_template" helper:"failureEmailTemplate,emptyIsNull,default:"`
	FailureEmailName            types.String `tfsdk:"failure_email_name" helper:"failureEmailName,emptyIsNull,default:"`
	SuccessEmailNotification    types.Bool   `tfsdk:"success_email_notification" helper:"successEmailNotification,emptyIsNull,default:"`
	SuccessEmailTemplate        types.String `tfsdk:"success_email_template" helper:"successEmailTemplate,emptyIsNull,default:"`
	SuccessEmailName            types.String `tfsdk:"success_email_name" helper:"successEmailName,emptyIsNull,default:"`
	TriggeringEmailNotification types.Bool   `tfsdk:"triggering_email_notification" helper:"triggeringEmailNotification,emptyIsNull,default:"`
	TriggeringEmailTemplate     types.String `tfsdk:"triggering_email_template" helper:"triggeringEmailTemplate,emptyIsNull,default:"`
	TriggeringEmailName         types.String `tfsdk:"triggering_email_name" helper:"triggeringEmailName,emptyIsNull,default:"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	/*StepStatuses                []struct {
		Id     types.String `tfsdk:"id"`
		StepId types.String `tfsdk:"step_id"`
		Status types.String `tfsdk:"status"`
	} `tfsdk:"step_statuses"`*/
	//Steps types.List `tfsdk:"steps" helper:"steps,fold,elementtype:stRouteStepSendToPartnerResourceModel,optional"`
	//Steps types.List `tfsdk:"steps" helper:"steps,fold:type,elementtype:stRouteSimpleSteps,optional"`
	Steps []stRouteSimpleSteps `tfsdk:"steps" helper:"steps,fold:type,optional"`
}

func NewSTRouteSimpleResource() resource.Resource {
	return NewSTResource(&stRouteSimpleResourceModel{}, "st_route_simple", "", "/api/v2.0/routes", "/api/v2.0/routes/{id}").IgnoreDeleteNotFoundError()
}

func init() {
	registerResource(NewSTRouteSimpleResource)
}
