package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stRouteCompositeResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Description   types.String `tfsdk:"description"`
	Type          types.String `tfsdk:"type" helper:",default:COMPOSITE"`
	ManagedByCG   types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,emptyIsNull,default:"`
	RouteTemplate types.String `tfsdk:"route_template" helper:"routeTemplate,required"`
	Account       types.String `tfsdk:"account" helper:",optional"`

	ConditionType               types.String `tfsdk:"condition_type" helper:"conditionType,default:MATCH_ALL"`
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

	Subscriptions types.List `tfsdk:"subscriptions" helper:"subscriptions,elementtype:string,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	/*StepStatuses                []struct {
		Id     types.String `tfsdk:"id"`
		StepId types.String `tfsdk:"step_id"`
		Status types.String `tfsdk:"status"`
	} `tfsdk:"step_statuses"`*/
	Steps []stRouteStepExecuteRouteResourceModel `tfsdk:"steps" helper:"steps,optional"`
}

func NewSTRouteCompositeResource() resource.Resource {
	return NewSTResource(&stRouteCompositeResourceModel{}, "st_route_composite", "", "/api/v2.0/routes", "/api/v2.0/routes/{id}")
}

func init() {
	registerResource(NewSTRouteCompositeResource)
}
