package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteHTTPModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                                     types.String `tfsdk:"type" helper:",default:http"`
	Protocol                                 types.String `tfsdk:"protocol" helper:",default:http"`
	TransferType                             types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:internal"`
	MaxConcurrentConnection                  types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                                  types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel                              types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                                  types.String `tfsdk:"account" helper:",required"`
	Host                                     types.String `tfsdk:"host" helper:",required"`
	Port                                     types.String `tfsdk:"port" helper:",default:22"`
	Dmz                                      types.String `tfsdk:"dmz" helper:",default:none"`
	DownloadFolderAdvancedExpressionEnabled  types.Bool   `tfsdk:"download_folder_advanced_expression_enabled" helper:"downloadFolderAdvancedExpressionEnabled,default"`
	DownloadFolder                           types.String `tfsdk:"download_folder" helper:"downloadFolder,default:/"`
	DownloadPatternAdvancedExpressionEnabled types.Bool   `tfsdk:"download_pattern_advanced_expression_enabled" helper:"downloadPatternAdvancedExpressionEnabled,default"`
	DownloadPattern                          types.String `tfsdk:"download_pattern" helper:"downloadPattern,default:*"`
	UploadFolderOverridable                  types.Bool   `tfsdk:"upload_folder_overridable" helper:"uploadFolderOverridable,default"`
	UploadFolder                             types.String `tfsdk:"upload_folder" helper:"uploadFolder,default"`
	TransferMode                             types.String `tfsdk:"transfer_mode" helper:"transferMode,enum:/AUTO_DETECT/BINARY/ASCII,default:AUTO_DETECT"`
	FipsMode                                 types.Bool   `tfsdk:"fips_mode" helper:"fipsMode,default:false"`
	Uri                                      types.String `tfsdk:"uri" helper:"uri,emptyIsNull,default:"`
	UseUri                                   types.Bool   `tfsdk:"use_uri" helper:"useUri,default:false"`
	UserName                                 types.String `tfsdk:"user_name" helper:"userName,required"`
	UsePassword                              types.Bool   `tfsdk:"use_password" helper:"usePassword,default:true"`
	UsePasswordExpr                          types.Bool   `tfsdk:"use_password_expr" helper:"usePasswordExpr,default:false"`
	Password                                 types.String `tfsdk:"password" helper:"password,default,noread"`
	ClientCertificate                        types.String `tfsdk:"client_certificate" helper:"clientCertificate,emptyIsNull,default:"`
	CipherSuites                             types.String `tfsdk:"cipher_suites" helper:"cipherSuites,computed,optional"`
	Protocols                                types.String `tfsdk:"protocols" helper:"protocols,emptyIsNull,computed,optional"`
	IsSecure                                 types.Bool   `tfsdk:"is_secure" helper:"isSecure,default:false"`
	RequestMode                              types.String `tfsdk:"request_mode" helper:"requestMode,enum:/GET/POST,default:GET"`
	VerifyCert                               types.Bool   `tfsdk:"verify_cert" helper:"verifyCert,default:false"`

	PostTransmissionActions struct {
		DeleteOnPermFailIn  types.Bool   `tfsdk:"delete_on_perm_fail_in" helper:"deleteOnPermFailIn,default:false"`
		DeleteOnPermFailOut types.Bool   `tfsdk:"delete_on_perm_fail_out" helper:"deleteOnPermFailOut,default:false"`
		DeleteOnSuccessIn   types.Bool   `tfsdk:"delete_on_success_in" helper:"deleteOnSuccessIn,default:false"`
		DeleteOnTempFailOut types.Bool   `tfsdk:"delete_on_temp_fail_out" helper:"deleteOnTempFailOut,default:false"`
		DoAsOut             types.String `tfsdk:"do_as_out" helper:"doAsOut,default"`
		DoAsIn              types.String `tfsdk:"do_as_in" helper:"doAsIn,default"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,computed,default:"`
	AlternativeAddresses []struct {
		Id       types.String `tfsdk:"id" helper:",computed,state"`
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteHTTPModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteHTTPModel{}, "st_site_http", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=http]")
}

func init() {
	registerResource(NewSTTransferSiteHTTPModelResource)
}
