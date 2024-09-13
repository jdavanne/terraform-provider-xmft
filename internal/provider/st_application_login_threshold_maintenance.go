package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stLoginThresholdMaintenanceApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Emails              types.String `tfsdk:"emails" helper:"emails,default:"`
	IsReport            types.Bool   `tfsdk:"is_report" helper:"isReport,default:false"`
	IsUnlock            types.Bool   `tfsdk:"is_unlock" helper:"isUnlock,default:false"`
	ManagedByCG         types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes               types.String `tfsdk:"notes" helper:"notes,default:"`
	ReportEmailTemplate types.String `tfsdk:"report_email_template" helper:"reportEmailTemplate,enum:/LoginThresholdReport.xhtml/AdhocDefault.xhtml/LoginThresholdNotification.xhtml/RoutingFailedNotification.xhtml/RoutingSucceededNotification.xhtml/RoutingTriggeredNotification.xhtml/AccountMaintenanceNotification.xhtml/FileMaintenanceNotification.xhtml,default:LoginThresholdReport.xhtml"`
	Type                types.String `tfsdk:"type" helper:"type,default:LoginThresholdMaintenance"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewSTLoginThresholdMaintenanceApplicationResource() resource.Resource {
	return NewSTResource(&stLoginThresholdMaintenanceApplication{}, "st_login_threshold_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=LoginThresholdMaintenance]")
}

func init() {
	registerResource(NewSTLoginThresholdMaintenanceApplicationResource)
}
