package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stPackageRetentionMaintApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	ManagedByCG       types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	MaxRunningMinutes types.Int64  `tfsdk:"max_running_minutes" helper:"maxRunningMinutes,default:"`
	Notes             types.String `tfsdk:"notes" helper:"notes,default:"`
	Type              types.String `tfsdk:"type" helper:"type,default:PackageRetentionMaint"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewSTPackageRetentionMaintApplicationResource() resource.Resource {
	return NewSTResource(&stPackageRetentionMaintApplication{}, "st_package_retention_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=PackageRetentionMaint]")
}

func init() {
	registerResource(NewSTPackageRetentionMaintApplicationResource)
}
