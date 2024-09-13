package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stHumanSystemApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	ManagedByCG types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes       types.String `tfsdk:"notes" helper:"notes,default:"`
	Type        types.String `tfsdk:"type" helper:"type,default:HumanSystem"`
}

func NewSTHumanSystemApplicationResource() resource.Resource {
	return NewSTResource(&stHumanSystemApplication{}, "st_human_system_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=HumanSystem]")
}

func init() {
	registerResource(NewSTHumanSystemApplicationResource)
}
