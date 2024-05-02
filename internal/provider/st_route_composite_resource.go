package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stRouteCompositeResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Description                 types.String   `tfsdk:"description"`
	Type                        types.String   `tfsdk:"type" helper:",default:COMPOSITE"`
	ManagedByCG                 types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,optional,computed"`
	RouteTemplate               types.String   `tfsdk:"route_template" helper:"routeTemplate,required"`
	Account                     types.String   `tfsdk:"account" helper:",optional"`
	Subscriptions               []types.String `tfsdk:"subscriptions"`
	ConditionType               types.String   `tfsdk:"condition_type" helper:"conditionType,default:MATCH_ALL"`
	Condition                   types.String   `tfsdk:"condition" helper:",optional,computed"`
	FailureEmailNotification    types.Bool     `tfsdk:"failure_email_notification" helper:"failureEmailNotification,optional,computed"`
	FailureEmailTemplate        types.String   `tfsdk:"failure_email_template" helper:"failureEmailTemplate,optional,computed"`
	FailureEmailName            types.String   `tfsdk:"failure_email_name" helper:"failureEmailName,optional,computed"`
	SuccessEmailNotification    types.Bool     `tfsdk:"success_email_notification" helper:"successEmailNotification,optional,computed"`
	SuccessEmailTemplate        types.String   `tfsdk:"success_email_template" helper:"successEmailTemplate,optional,computed"`
	SuccessEmailName            types.String   `tfsdk:"success_email_name" helper:"successEmailName,optional,computed"`
	TriggeringEmailNotification types.Bool     `tfsdk:"triggering_email_notification" helper:"triggeringEmailNotification,optional,computed"`
	TriggeringEmailTemplate     types.String   `tfsdk:"triggering_email_template" helper:"triggeringEmailTemplate,optional,computed"`
	TriggeringEmailName         types.String   `tfsdk:"triggering_email_name" helper:"triggeringEmailName,optional,computed"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	/*StepStatuses                []struct {
		Id     types.String `tfsdk:"id"`
		StepId types.String `tfsdk:"step_id"`
		Status types.String `tfsdk:"status"`
	} `tfsdk:"step_statuses"`*/
	Steps types.List `tfsdk:"steps" helper:"steps,elementtype:stRouteStepExecuteRouteResourceModel,optional"`
}

func NewSTRouteCompositeResource() resource.Resource {
	return NewSTResource(&stRouteCompositeResourceModel{}, "st_route_composite", "", "/api/v2.0/routes", "/api/v2.0/routes/{id}")
}
