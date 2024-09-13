package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stAuditLogMaintApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	DayMode       types.Int64  `tfsdk:"day_mode" helper:"dayMode,default:"`
	ExportDir     types.String `tfsdk:"export_dir" helper:"exportDir,default:"`
	ExportEnabled types.Bool   `tfsdk:"export_enabled" helper:"exportEnabled,default:false"`

	ManagedByCG         types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes               types.String `tfsdk:"notes" helper:"notes,default:"`
	TimeKeepingDataInDb types.Int64  `tfsdk:"time_keeping_data_in_db" helper:"timeKeepingDataInDb,default:"`
	Type                types.String `tfsdk:"type" helper:"type,default:AuditLogMaint"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewAuditLogMaintApplicationResource() resource.Resource {
	return NewSTResource(&stAuditLogMaintApplication{}, "st_audit_log_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=AuditLogMaint]")
}

func init() {
	registerResource(NewAuditLogMaintApplicationResource)
}
