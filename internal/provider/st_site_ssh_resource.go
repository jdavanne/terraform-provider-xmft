package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteSSHModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                                     types.String `tfsdk:"type" helper:",default:ssh"`
	Protocol                                 types.String `tfsdk:"protocol" helper:",default:ssh"`
	TransferType                             types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection                  types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                                  types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel                              types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                                  types.String `tfsdk:"account" helper:",required"`
	Host                                     types.String `tfsdk:"host" helper:",required"`
	Port                                     types.String `tfsdk:"port" helper:",default:22"`
	Dmz                                      types.String `tfsdk:"dmz" helper:",default:none"`
	DownloadFolderAdvancedExpressionEnabled  types.Bool   `tfsdk:"download_folder_advanced_expression_enabled" helper:"downloadFolderAdvancedExpressionEnabled,default"`
	DownloadFolder                           types.String `tfsdk:"download_folder" helper:"downloadFolder,default:"`
	DownloadPatternAdvancedExpressionEnabled types.Bool   `tfsdk:"download_pattern_advanced_expression_enabled" helper:"downloadPatternAdvancedExpressionEnabled,default"`
	DownloadPatternType                      types.String `tfsdk:"download_pattern_type" helper:"downloadPatternType,default:glob"`
	DownloadPattern                          types.String `tfsdk:"download_pattern" helper:"downloadPattern,default:*"`
	UploadFolderOverridable                  types.Bool   `tfsdk:"upload_folder_overridable" helper:"uploadFolderOverridable,default"`
	UploadFolder                             types.String `tfsdk:"upload_folder" helper:"uploadFolder,default"`
	UploadPermissions                        types.String `tfsdk:"upload_permissions" helper:"uploadPermissions,default:0644"`
	UpdatePermissionsWithChmodCommand        types.String `tfsdk:"update_permissions_with_chmod_command" helper:"updatePermissionsWithChmodCommand,emptyIsNull,default:"`
	TransferMode                             types.String `tfsdk:"transfer_mode" helper:"transferMode,default:AUTO_DETECT"`
	VerifyFinger                             types.Bool   `tfsdk:"verify_finger" helper:"verifyFinger,default:false"`
	FingerPrint                              types.String `tfsdk:"finger_print" helper:"fingerPrint,default"`
	FipsMode                                 types.Bool   `tfsdk:"fips_mode" helper:"fipsMode,default:false"`
	UserName                                 types.String `tfsdk:"user_name" helper:"userName,required"`
	UsePassword                              types.Bool   `tfsdk:"use_password" helper:"usePassword,default:true"`
	UsePasswordExpr                          types.Bool   `tfsdk:"use_password_expr" helper:"usePasswordExpr,default:false"`
	Password                                 types.String `tfsdk:"password" helper:"password,default,noread"`
	SocketTimeout                            types.Int64  `tfsdk:"socket_timeout" helper:"socketTimeout,default:300"`
	SocketBufferSize                         types.Int64  `tfsdk:"socket_buffer_size" helper:"socketBufferSize,default:65536"`
	SocketSendBufferSize                     types.Int64  `tfsdk:"socket_send_buffer_size" helper:"socketSendBufferSize,default:65536"`
	BufferSize                               types.Int64  `tfsdk:"buffer_size" helper:"bufferSize,default:32768"`
	BlockSize                                types.Int64  `tfsdk:"block_size" helper:"blockSize,default:32768"`
	TcpNoDelay                               types.Bool   `tfsdk:"tcp_no_delay" helper:"tcpNoDelay,default:true"`
	ClientCertificate                        types.String `tfsdk:"client_certificate" helper:"clientCertificate,emptyIsNull,default:"`
	ConnectionIdleTimeout                    types.Int64  `tfsdk:"connection_idle_timeout" helper:"connectionIdleTimeout,default:300"`
	CipherSuites                             types.String `tfsdk:"cipher_suites" helper:"cipherSuites,computed,optional"`
	Protocols                                types.String `tfsdk:"protocols" helper:"protocols,default:"`
	AllowedMacs                              types.String `tfsdk:"allowed_macs" helper:"allowedMacs,computed,optional"`
	KeyExchangeAlgorithms                    types.String `tfsdk:"key_exchange_algorithms" helper:"keyExchangeAlgorithms,computed,optional"`
	PublicKeys                               types.String `tfsdk:"public_keys" helper:"publicKeys,computed,optional"`

	PostTransmissionActions struct {
		DeleteOnTempFailOut types.Bool   `tfsdk:"delete_on_temp_fail_out" helper:"deleteOnTempFailOut,default:false"`
		DeleteOnSuccessIn   types.Bool   `tfsdk:"delete_on_success_in" helper:"deleteOnSuccessIn,default:false"`
		DeleteOnPermFailOut types.Bool   `tfsdk:"delete_on_perm_fail_out" helper:"deleteOnPermFailOut,default:false"`
		DeleteOnPermFailIn  types.Bool   `tfsdk:"delete_on_perm_fail_in" helper:"deleteOnPermFailIn,default:false"`
		DoAsOut             types.String `tfsdk:"do_as_out" helper:"doAsOut,emptyIsNull,default"`
		DoAsIn              types.String `tfsdk:"do_as_in" helper:"doAsIn,emptyIsNull,default"`
		MoveOnTempFailOut   types.String `tfsdk:"move_on_temp_fail_out" helper:"moveOnTempFailOut,emptyIsNull,default"`
		MoveOnPermFailOut   types.String `tfsdk:"move_on_perm_fail_out" helper:"moveOnPermFailOut,emptyIsNull,default"`
		MoveOnSuccessOut    types.String `tfsdk:"move_on_success_out" helper:"moveOnSuccessOut,emptyIsNull,default"`
		MoveOnPermFailIn    types.String `tfsdk:"move_on_perm_fail_in" helper:"moveOnPermFailIn,emptyIsNull,default"`
		MoveOnSuccessIn     types.String `tfsdk:"move_on_success_in" helper:"moveOnSuccessIn,emptyIsNull,default"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,computed,default:"`

	AlternativeAddresses []struct {
		Id       types.String `tfsdk:"id" helper:",computed,state"`
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteSSHModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteSSHModel{}, "st_site_ssh", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=ssh]")
}

func init() {
	registerResource(NewSTTransferSiteSSHModelResource)
}
