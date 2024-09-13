package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stUnlicensedAccountMaintApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	InactivePeriod types.Int64  `tfsdk:"inactive_period" helper:"inactivePeriod,default:"`
	ManagedByCG    types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes          types.String `tfsdk:"notes" helper:"notes,default:"`
	Type           types.String `tfsdk:"type" helper:"type,default:UnlicensedAccountMaint"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewSTUnlicensedAccountMaintApplicationResource() resource.Resource {
	return NewSTResource(&stUnlicensedAccountMaintApplication{}, "st_unlicensed_account_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=UnlicensedAccountMaint]")
}

func init() {
	registerResource(NewSTUnlicensedAccountMaintApplicationResource)
}
