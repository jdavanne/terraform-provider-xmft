package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteSynchronyTransferModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	AccessLevel          types.String `tfsdk:"access_level" helper:"accessLevel,enum:/PRIVATE/PUBLIC/BUSINESS_UNIT,default:"`
	Account              types.String `tfsdk:"account" helper:"account,default:"`
	AdditionalProperties struct {
		Appcycid  types.String `tfsdk:"appcycid" helper:"appcycid,default:"`
		Appobjid  types.String `tfsdk:"appobjid" helper:"appobjid,default:"`
		Direction types.String `tfsdk:"direction" helper:"direction,default:"`
		FName     types.String `tfsdk:"f_name" helper:"fName,default:"`
		FreeMsg   types.String `tfsdk:"free_msg" helper:"freeMsg,default:"`
		Ida       types.String `tfsdk:"ida" helper:"ida,default:"`
		Ipart     types.String `tfsdk:"ipart" helper:"ipart,default:"`
		MaxDate   types.String `tfsdk:"max_date" helper:"maxDate,default:"`
		MaxTime   types.String `tfsdk:"max_time" helper:"maxTime,default:"`
		MinDate   types.String `tfsdk:"min_date" helper:"minDate,default:"`
		MinTime   types.String `tfsdk:"min_time" helper:"minTime,default:"`
		NIdf      types.String `tfsdk:"n_idf" helper:"nIdf,default:"`
		NfName    types.String `tfsdk:"nf_name" helper:"nfName,default:"`
		Pri       types.String `tfsdk:"pri" helper:"pri,default:"`
		RUser     types.String `tfsdk:"r_user" helper:"rUser,default:"`
		Rappl     types.String `tfsdk:"rappl" helper:"rappl,default:"`
		SUser     types.String `tfsdk:"s_user" helper:"sUser,default:"`
		Sappl     types.String `tfsdk:"sappl" helper:"sappl,default:"`
		State     types.String `tfsdk:"state" helper:",default:"`
		Trk       types.String `tfsdk:"trk" helper:"trk,default:"`
	} `tfsdk:"additional_properties" helper:"additionalProperties,default:"`

	Default types.Bool   `tfsdk:"default" helper:",default:false"`
	Host    types.String `tfsdk:"host" helper:"host,default:"`

	MaxConcurrentConnection types.Int64 `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`

	Partner         types.String `tfsdk:"partner" helper:"partner,default:"`
	Password        types.String `tfsdk:"password" helper:"password,default:"`
	Port            types.String `tfsdk:"port" helper:"port,default:"`
	Protocol        types.String `tfsdk:"protocol" helper:"protocol,default:synchrony transfer"`
	SiteTemplate    types.String `tfsdk:"site_template" helper:"siteTemplate,default:"`
	TransferProfile types.String `tfsdk:"transfer_profile" helper:"transferProfile,default:"`
	TransferType    types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:false"`
	Type            types.String `tfsdk:"type" helper:"type,default:synchrony transfer"`
	UsePasswordExpr types.Bool   `tfsdk:"use_password_expr" helper:"usePasswordExpr,default:false"`
	Username        types.String `tfsdk:"username" helper:"username,default:"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteSynchronyTransferModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteSynchronyTransferModel{}, "st_site_synchrony_transfer", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=synchrony transfer]")
}

func init() {
	registerResource(NewSTTransferSiteSynchronyTransferModelResource)
}
