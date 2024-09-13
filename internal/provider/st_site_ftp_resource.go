package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
	{
	  "type": "ftp",
	  "id": "2cd08585919ffdba0191e87cf5e2129c",
	  "name": "ftp1",
	  "account": "account1-2024-09-12_15-41-41-1726148501632479000",
	  "protocol": "ftp",
	  "transferType": "unspecified",
	  "maxConcurrentConnection": 0,
	  "default": false,
	  "accessLevel": "PRIVATE",
	  "isSecure": true,
	  "verifyCert": false,
	  "ccc": false,
	  "tlsShutdownCcc": false,
	  "host": "dsffdqd",
	  "port": "21",
	  "alternativeAddresses": [],
	  "uploadCommand": "STOR",
	  "downloadFolder": "/",
	  "downloadFolderAdvancedExpressionEnabled": false,
	  "downloadPattern": "*",
	  "downloadPatternAdvancedExpressionEnabled": false,
	  "uploadFolder": "/",
	  "uploadFolderOverridable": false,
	  "preferredMethod": "default",
	  "userName": "ddsfs",
	  "usePassword": true,
	  "usePasswordExpr": null,
	  "password": "{AES128}zCWDVVRrOthyPuc9dPdUfw==",
	  "clientCertificate": null,
	  "transferMode": "AUTO_DETECT",
	  "siteCommand": "",
	  "activeMode": false,
	  "fipsMode": false,
	  "dmz": "none",
	  "matchAnyLineTerminators": null,
	  "hostnameVerification": "DEFAULT",
	  "cipherSuites": "TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_CHACHA20_POLY1305_SHA256,TLS_AES_128_CCM_SHA256,TLS_AES_128_CCM_8_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_DHE_RSA_WITH_AES_256_GCM_SHA384,TLS_DHE_DSS_WITH_AES_256_GCM_SHA384,TLS_DHE_DSS_WITH_AES_256_CBC_SHA256,TLS_DHE_RSA_WITH_AES_256_CBC_SHA256,TLS_DHE_RSA_WITH_AES_128_GCM_SHA256,TLS_DHE_DSS_WITH_AES_128_GCM_SHA256,TLS_DHE_DSS_WITH_AES_128_CBC_SHA256,TLS_DHE_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA256",
	  "fipsCipherSuites": "",
	  "protocols": "TLSv1.2,TLSv1.3",
	  "postTransmissionActions": {
	    "doAsOut": null,
	    "deleteOnTempFailOut": null,
	    "moveOnTempFailOut": null,
	    "deleteOnPermFailOut": null,
	    "moveOnPermFailOut": null,
	    "moveOnSuccessOut": null,
	    "doAsIn": null,
	    "deleteOnPermFailIn": null,
	    "moveOnPermFailIn": null,
	    "deleteOnSuccessIn": null,
	    "moveOnSuccessIn": null
	  },
	  "additionalAttributes": {},
	  "metadata": {
	    "links": {
	      "account": "https://ci8.jda.axwaytest.net:8444/api/v2.0/accounts/account1-2024-09-12_15-41-41-1726148501632479000"
	    }
	  }
	}
*/
type stTransferSiteFTPModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                                     types.String `tfsdk:"type" helper:",default:ftp"`
	Protocol                                 types.String `tfsdk:"protocol" helper:",default:ftp"`
	TransferType                             types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection                  types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                                  types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel                              types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	ActiveMode                               types.Bool   `tfsdk:"active_mode" helper:"activeMode,default:false"`
	Account                                  types.String `tfsdk:"account" helper:",required"`
	CCC                                      types.Bool   `tfsdk:"ccc" helper:",emptyIsNull,default:false"`
	Host                                     types.String `tfsdk:"host" helper:",required"`
	Port                                     types.String `tfsdk:"port" helper:",default:22"`
	Dmz                                      types.String `tfsdk:"dmz" helper:",default:none"`
	DownloadFolderAdvancedExpressionEnabled  types.Bool   `tfsdk:"download_folder_advanced_expression_enabled" helper:"downloadFolderAdvancedExpressionEnabled,default:false"`
	DownloadFolder                           types.String `tfsdk:"download_folder" helper:"downloadFolder,default:/"`
	DownloadPatternAdvancedExpressionEnabled types.Bool   `tfsdk:"download_pattern_advanced_expression_enabled" helper:"downloadPatternAdvancedExpressionEnabled,default:false"`
	DownloadPattern                          types.String `tfsdk:"download_pattern" helper:"downloadPattern,default:*"`
	TransferMode                             types.String `tfsdk:"transfer_mode" helper:"transferMode,default:AUTO_DETECT"`
	Password                                 types.String `tfsdk:"password" helper:"password,default,noread"`
	FipsMode                                 types.Bool   `tfsdk:"fips_mode" helper:"fipsMode,default:false"`
	ClientCertificate                        types.String `tfsdk:"client_certificate" helper:"clientCertificate,emptyIsNull,default:"`
	CipherSuites                             types.String `tfsdk:"cipher_suites" helper:"cipherSuites,computed,optional"`
	HostnameVerification                     types.String `tfsdk:"hostname_verification" helper:"hostnameVerification,emptyIsNull,default:DEFAULT"`
	IsSecure                                 types.Bool   `tfsdk:"is_secure" helper:"isSecure,default:false"`
	MatchAnyLineTerminators                  types.Bool   `tfsdk:"match_any_line_terminators" helper:"matchAnyLineTerminators,emptyIsNull,default:false"`
	PreferredMethod                          types.String `tfsdk:"preferred_method" helper:"preferredMethod,emptyIsNull,default:default"`
	Protocols                                types.String `tfsdk:"protocols" helper:"protocols,emptyIsNull,computed,optional"`
	SiteCommand                              types.String `tfsdk:"site_command" helper:"siteCommand,default:"`
	TlsShutDownCcc                           types.Bool   `tfsdk:"tls_shutdown_ccc" helper:"tlsShutdownCcc,emptyIsNull,default:false"`
	UploadCommand                            types.String `tfsdk:"upload_command" helper:"uploadCommand,emptyIsNull,default:STOR"`
	UploadFolderOverridable                  types.Bool   `tfsdk:"upload_folder_overridable" helper:"uploadFolderOverridable,default:false"`
	UploadFolder                             types.String `tfsdk:"upload_folder" helper:"uploadFolder,emptyIsNull,default:/"`
	UserName                                 types.String `tfsdk:"user_name" helper:"userName,required"`
	UsePassword                              types.Bool   `tfsdk:"use_password" helper:"usePassword,default:false"`
	UsePasswordExpr                          types.Bool   `tfsdk:"use_password_expr" helper:"usePasswordExpr,emptyIsNull,default:"`
	VerifyCert                               types.Bool   `tfsdk:"verify_cert" helper:"verifyCert,default:false"`

	PostTransmissionActions struct {
		DeleteOnPermFailIn  types.Bool   `tfsdk:"delete_on_perm_fail_in" helper:"deleteOnPermFailIn,emptyIsNull,default:false"`
		DeleteOnPermFailOut types.Bool   `tfsdk:"delete_on_perm_fail_out" helper:"deleteOnPermFailOut,emptyIsNull,default:false"`
		DeleteOnSuccessIn   types.Bool   `tfsdk:"delete_on_success_in" helper:"deleteOnSuccessIn,emptyIsNull,default:false"`
		DeleteOnTempFailOut types.Bool   `tfsdk:"delete_on_temp_fail_out" helper:"deleteOnTempFailOut,emptyIsNull,default:false"`
		MoveOnTempFailOut   types.String `tfsdk:"move_on_temp_fail_out" helper:"moveOnTempFailOut,emptyIsNull,default"`
		MoveOnPermFailOut   types.String `tfsdk:"move_on_perm_fail_out" helper:"moveOnPermFailOut,emptyIsNull,default"`
		MoveOnSuccessOut    types.String `tfsdk:"move_on_success_out" helper:"moveOnSuccessOut,emptyIsNull,default"`
		MoveOnPermFailIn    types.String `tfsdk:"move_on_perm_fail_in" helper:"moveOnPermFailIn,emptyIsNull,default"`
		MoveOnSuccessIn     types.String `tfsdk:"move_on_success_in" helper:"moveOnSuccessIn,emptyIsNull,default"`
		DoAsOut             types.String `tfsdk:"do_as_out" helper:"doAsOut,emptyIsNull,default"`
		DoAsIn              types.String `tfsdk:"do_as_in" helper:"doAsIn,emptyIsNull,default"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,computed,default:"`

	AlternativeAddresses []struct {
		Id       types.String `tfsdk:"id" helper:",computed,state"`
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteFTPModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteFTPModel{}, "st_site_ftp", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=ftp]")
}

func init() {
	registerResource(NewSTTransferSiteFTPModelResource)
}
