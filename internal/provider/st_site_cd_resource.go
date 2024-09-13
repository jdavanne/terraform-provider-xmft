package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteCDModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                    types.String `tfsdk:"type" helper:",default:cd"`
	Protocol                types.String `tfsdk:"protocol" helper:",default:cd"`
	TransferType            types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                 types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel             types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                 types.String `tfsdk:"account" helper:",required"`
	TransferMode            types.String `tfsdk:"transfer_mode" helper:"transferMode,default:AUTO_DETECT"`
	UserName                types.String `tfsdk:"user_name" helper:"userName,required"`
	UsePassword             types.Bool   `tfsdk:"use_password" helper:"usePassword,default:true"`
	UsePasswordExpr         types.Bool   `tfsdk:"use_password_expr" helper:"usePasswordExpr,default:false"`
	Password                types.String `tfsdk:"password" helper:"password,default,noread"`
	CertificateId           types.String `tfsdk:"certificate_id" helper:"certificateId,default:"`
	UseCertificate          types.Bool   `tfsdk:"use_certificate" helper:"useCertificate,default:false"`
	LocalServerName         types.String `tfsdk:"local_server_name" helper:"localServerName,default:"`
	LocalServerPort         types.String `tfsdk:"local_server_port" helper:"localServerPort,default:"`
	RecScript               types.String `tfsdk:"rec_script" helper:"recScript,default:"`
	SendScript              types.String `tfsdk:"send_script" helper:"sendScript,default:"`
	SiteTemplate            types.String `tfsdk:"site_template" helper:"siteTemplate,default:"`
	SubmitScriptWithinFile  types.Bool   `tfsdk:"submit_script_within_file" helper:"submitScriptWithinFile,default:false"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteCDModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteCDModel{}, "st_site_cd", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=cd]")
}

func init() {
	registerResource(NewSTTransferSiteCDModelResource)
}
