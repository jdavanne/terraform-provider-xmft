package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stMBFTApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	ManagedByCG types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes       types.String `tfsdk:"notes" helper:"notes,default:"`
	Type        types.String `tfsdk:"type" helper:"type,default:MBFT"`
}

func NewSTMBFTApplicationResource() resource.Resource {
	return NewSTResource(&stMBFTApplication{}, "st_mbft_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=MBFT]")
}

func init() {
	registerResource(NewSTMBFTApplicationResource)
}
