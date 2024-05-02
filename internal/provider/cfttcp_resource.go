package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cftTcpResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required,noread,nowrite"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Cnxin    types.String `tfsdk:"cnxin"`
	Cnxout   types.String `tfsdk:"cnxout"`
	Cnxinout types.String `tfsdk:"cnxinout"`

	Omaxtime types.String `tfsdk:"omaxtime"  helper:",computed,optional"`
	Omintime types.String `tfsdk:"omintime"  helper:",computed,optional"`
	Imaxtime types.String `tfsdk:"imaxtime"  helper:",computed,optional"`
	Imintime types.String `tfsdk:"imintime"  helper:",computed,optional"`

	Retrym types.String `tfsdk:"retrym"`
	Retryn types.String `tfsdk:"retryn"`
	Retryw types.String `tfsdk:"retryw"`

	Host types.String `tfsdk:"host" helper:",computed,optional"`
}

func NewCFTTcpResource() resource.Resource {
	return NewCFTResource(&cftTcpResourceModel{}, "cfttcp", "cfttcp", "/cft/api/v1/objects/cftpart/{name}/tcp", "/cft/api/v1/objects/cftpart/{name}/tcp/1")
}
